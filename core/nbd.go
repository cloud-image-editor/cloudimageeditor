package cloudimageeditor

/*
 * cloudimageeditor package
 *
 * Authors:
 *     zhenwei pi <cloudimageeditor@126.com>
 *
 * This project is under Apache v2 License.
 *
 */

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	cloudimageeditorutil "github.com/cloud-image-editor/cloudimageeditor/util"
)

const (
	// NbdModulePath : nbd kmod parameters folder path
	NbdModulePath = "/sys/module/nbd/parameters/"
	// NbdMaxPart : nbd kmod max_part path
	NbdMaxPart = NbdModulePath + "max_part"
)

// chechNbd : check kmod parameters, we expect NbdMaxPart is at least 8. otherwise reload nbd kmod by rmmod & modprobe
func checkNbdKmod() error {
	maxPart, err := cloudimageeditorutil.GetProcInt64(NbdMaxPart)
	if err == nil {
		if maxPart >= 8 {
			return nil
		}

		_, err := cloudimageeditorutil.Execute("rmmod", []string{"nbd"})
		if err != nil {
			return err
		}
	}

	_, err = cloudimageeditorutil.Execute("modprobe", []string{"nbd", "max_part=8", "nbds_max=16"})

	return err
}

// getFreeNbdDev : try to scan all /dev/nbdX device until find a unused one.
//                 if succeed, return @dev name like "/dev/nbd0"
func getFreeNbdDev() (dev string, err error) {
	index := 0

	for index < 16 {
		name := fmt.Sprintf("/dev/nbd%d", index)
		f, err := os.Open(name)
		if err != nil {
			return "", err
		}
		defer f.Close()
		index++

		head := make([]byte, 512)
		headn, err := f.Read(head)
		if err != nil {
			if err == io.EOF {
				return name, nil
			}
			cloudimageeditorutil.LogD("getFreeNbdDev | read %s head error:%v", name, err)
			continue
		}

		if headn == 512 {
			// read nbd device first sector success, it should be in use
			continue
		} else if headn == 0 {
			return name, nil
		} else {
			cloudimageeditorutil.LogD("getFreeNbdDev | read %s first sector with %d bytes, unknown status, skip", name, headn)
			continue
		}
	}

	return "", fmt.Errorf("nbd device all in use")
}

// getMountPath : generate nbd device mount path
func getMountPath(nbdDev string) string {
	return "/var/cloudimageeditor" + nbdDev
}

// checkFolderMounted : check a folder is mounted or not
func checkFolderMounted(path string) (bool, error) {
	mounted := false
	mountTable := "/proc/mounts"

	mountInfo, err := ioutil.ReadFile(mountTable)
	if err != nil {
		return mounted, err
	}

	mountInfoLines := strings.Split(string(mountInfo), "\n")
	// analysis lines in mountInfo, Ex, /dev/nbd0p1 /var/cloudimageeditor/dev/nbd0 ext4 rw,relatime,data=ordered 0 0
	for _, line := range mountInfoLines {
		mounts := strings.Split(line, " ")
		if len(mounts) < 2 {
			continue // maybe a BUG
		}
		if path == mounts[1] {
			mounted = true // bingo
			break
		}
	}

	return mounted, nil
}

// MountImageByNbd : we can access @imgPath by qemu-nbd. if succeed, we can get @nbdDev & @mntPath
//                   why qemu-nbd ? qemu-nbd can recognize and opearte raw/qcow2 type by itself.
func MountImageByNbd(imgPath string) (nbdDev string, mntPath string, err error) {
	err = checkNbdKmod()
	if err != nil {
		return "", "", err
	}

	nbdDev, err = getFreeNbdDev()
	if err != nil {
		return "", "", err
	}
	cloudimageeditorutil.LogD("mountImageByNbd|getFreeNbdDev:%s", nbdDev)

	mntPath = getMountPath(nbdDev)
	_, err = os.Stat(mntPath)
	if err == nil {
		mounted, err := checkFolderMounted(mntPath)
		if err != nil {
			return "", "", fmt.Errorf("check /proc/mounts failed, " + err.Error())
		}
		if mounted {
			return "", "", fmt.Errorf("mntPath is mounted, maybe a bug")
		}
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(mntPath, os.ModeDir|0700); err != nil {
			return "", "", err
		}
	}

	result, err := cloudimageeditorutil.Execute("qemu-nbd", []string{"-c", nbdDev, imgPath})
	if err != nil {
		if strings.Contains(result, "WARNING") {
			// qemu-nbd auto detect format, if raw, qemu may report WARNING. BUT, it's NOT a error
			cloudimageeditorutil.LogD("MountImageByNbd | qemu-nbd report WARNING, not a fatal error\n")
		} else {
			return nbdDev, "", fmt.Errorf("qemu-nbd connect failed, %s, output %s", err.Error(), result)
		}
	}

	_, err = cloudimageeditorutil.Execute("partprobe", []string{nbdDev})
	if err != nil {
		time.Sleep(time.Second * 2)
	}

	// check nbd device first partation, and mount it to folder
	nbdDevp1 := nbdDev + "p1"
	_, err = os.Stat(nbdDevp1)
	if err != nil {
		return "", "", fmt.Errorf("device first partation not found")
	}

	result, err = cloudimageeditorutil.Execute("mount", []string{"-o", "rw,sync", nbdDevp1, mntPath})
	if err != nil {
		return "", "", fmt.Errorf("mount partation failed, %s, mount output : %s", err.Error(), result)
	}

	return nbdDev, mntPath, nil
}

// UnmountImageFreeNbd : unmount image and free nbd device
func UnmountImageFreeNbd(nbdDev string, mntPath string) (err error) {
	// try sync all mntPath
	cloudimageeditorutil.Execute("sync", []string{mntPath})

	// if umount fail, it's usually a network error. and this routine should be blocked.
	_, err = cloudimageeditorutil.Execute("umount", []string{mntPath})
	if err != nil {
		cloudimageeditorutil.LogD("unmountImageFreeNbd | umount %s, err:%s", mntPath, err.Error())
	}

	_, err = cloudimageeditorutil.Execute("qemu-nbd", []string{"-d", nbdDev})
	if err != nil {
		cloudimageeditorutil.LogD("umountVMFreeNbd | qemu-nbd -d %s fail, err:%s", nbdDev, err.Error())
	}

	os.RemoveAll(mntPath)

	return err
}
