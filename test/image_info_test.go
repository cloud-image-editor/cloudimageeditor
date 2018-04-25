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
	"testing"

	cloudimageeditorcore "github.com/cloud-image-editor/cloudimageeditor/core"
)

func TestImageInfo(t *testing.T) {
	//imagePath := "ubuntu.qcow2"
	//imagePath := "iscsi://192.168.9.218:20079/iqn.2017-12.vemerge.cn:zgvmyxvsdcndt01fvcnkzwzhdwx0i2fsd2f5cy10zxn0lwnvbnrhaw5lci12bs1zzge243142724/1"
	imagePath := "http://cloud.centos.org/centos/7/images/CentOS-7-x86_64-GenericCloud.qcow2"
	imgFormat, virtualSize, diskSize, err := cloudimageeditorcore.ImageInfo(imagePath)
	if err != nil {
		t.Fatalf("cloudimageeditorcore.ImageInfo : %s\n", err.Error())
	}

	t.Logf("IMAGE %s : format = %v, virtual size = %v, disk size = %v\n", imagePath, imgFormat, virtualSize, diskSize)
}
