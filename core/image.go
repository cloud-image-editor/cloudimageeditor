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
	"strings"

	cloudimageeditorutil "github.com/cloud-image-editor/cloudimageeditor/util"
)

// imageInfo : get @imgPath info.
//             maybe https://github.com/zchee/go-qcow2.git is a good choice. but this package only support local file.
//             since we need get http and iscsi file and so on.
func imageInfo(imgPath string) (imgFormat string, virtualSize uint64, diskSize uint64, err error) {
	result, err := cloudimageeditorutil.Execute("qemu-img", []string{"info", imgPath})
	if err != nil {
		return "", 0, 0, fmt.Errorf("qemu-img info failed, %s, result : %s", err.Error(), result)
	}
	// get image info by qemu-img info xxx, result like this :
	// 		file format: raw
	// 		virtual size: 10G (10737418240 bytes)
	// 		disk size: unavailable
	for _, line := range strings.Split(result, "\n") {
		if strings.HasPrefix(line, "file format:") {
			imgFormat = line[len("file format:"):len(line)]
			imgFormat = cloudimageeditorutil.StringKeepAlphaAndNum(imgFormat)
		} else if strings.HasPrefix(line, "virtual size:") {
			if strings.Contains(line, "(") && strings.Contains(line, ")") {
				vSizeString := line[strings.Index(line, "(")+1 : strings.Index(line, ")")]
				fmt.Sscanf(vSizeString, "%d bytes", &virtualSize)
			} else {
				virtualSize = 0
			}
		} else if strings.HasPrefix(line, "disk size:") {
			if strings.Contains(line, "unavailable") {
				diskSize = 0
			} else {
				dSizeString := line[len("disk size:"):len(line)]
				unitString := cloudimageeditorutil.StringKeepAlpha(dSizeString)
				sizeString := cloudimageeditorutil.StringKeepFloatNumber(dSizeString)
				diskSize, err = cloudimageeditorutil.ParseSizeWithUnit(sizeString, unitString)
				if err != nil {
					return "", 0, 0, err
				}
			}
		}
	}

	return imgFormat, virtualSize, diskSize, nil
}

// ImageInfo : get @imgPath info.
func ImageInfo(imgPath string) (imgFormat string, virtualSize uint64, diskSize uint64, err error) {
	return imageInfo(imgPath)
}

// ImageImport : import a image to another image. @srcFormat & @destFormat support "qcow2", "raw" and "auto".
//               "auto" means this API will try to guess.
//               @srcPath & @destPath should be already existed.
func ImageImport(srcPath string, srcFormat string, destPath string, destFormat string) error {
	sfmt, svSize, _, err := imageInfo(srcPath)
	if err != nil {
		return err
	}

	if (srcFormat != "auto") && (srcFormat != sfmt) {
		return fmt.Errorf("src format missmatch, parameter %s VS real %s", srcFormat, sfmt)
	}

	dfmt, dvSize, _, err := imageInfo(destPath)
	if err != nil {
		return err
	}

	if (destFormat != "auto") && (destFormat != sfmt) {
		return fmt.Errorf("dest format missmatch, parameter %s VS real %s", destFormat, dfmt)
	}

	if svSize > dvSize {
		return fmt.Errorf("dest image size %d is less than src image size %d", dvSize, svSize)
	}

	result, err := cloudimageeditorutil.Execute("qemu-img", []string{"convert", "-f", sfmt, "-O", dfmt, "-t", "none", srcPath, destPath})
	if err != nil {
		return fmt.Errorf("qemu-img convert failed, %s, result : %s", err.Error(), result)
	}

	return nil
}
