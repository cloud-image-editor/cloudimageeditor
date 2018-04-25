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
	"testing"

	cloudimageeditorcore "github.com/cloud-image-editor/cloudimageeditor/core"
	cloudimageeditormodule "github.com/cloud-image-editor/cloudimageeditor/module"
)

func TestFolder(t *testing.T) {
	cloudimageeditormodule.MuduleRegisterAll()

	imagePath := "ubuntu.qcow2"
	nbdDev, mntRoot, err := cloudimageeditorcore.MountImageByNbd(imagePath)
	if err != nil {
		t.Fatalf("cloudimageeditorcore.MountImageByNbd : %s\n", err.Error())
	}
	defer cloudimageeditorcore.UnmountImageFreeNbd(nbdDev, mntRoot)

	fmt.Printf("TestFolder : nbdDev = %s, mntRoot = %s\n", nbdDev, mntRoot)

	// in this case, folder "/root/sdb/redis" on host
	//		/root/sdb/redis
	//		/root/sdb/redis/bin/redis-server
	//		/root/sdb/redis/etc/redis.conf
	// after running this test case, we should see the same file in guest image.
	// we can use dpkg-deb -x xxx.deb yyy to unpack package, and inject them into guest.
	ifaceConf := make(map[string]string)
	ifaceConf["src"] = "redis"
	ifaceConf["dest"] = "/"

	err = cloudimageeditorcore.ExcuteModuleByName(mntRoot, "folder", ifaceConf)
	if err != nil {
		t.Fatalf("cloudimageeditorcore.ExcuteModuleByName : %s\n", err.Error())
	}
}
