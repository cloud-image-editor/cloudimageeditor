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
	cloudimageeditormodule "github.com/cloud-image-editor/cloudimageeditor/module"
)

func TestHostname(t *testing.T) {
	cloudimageeditormodule.MuduleRegisterAll()

	imagePath := "ubuntu.qcow2"
	//imagePath := "iscsi://192.168.9.218:20226/iqn.2017-12.vemerge.cn:testiscsivolume/1"
	nbdDev, mntRoot, err := cloudimageeditorcore.MountImageByNbd(imagePath)
	if err != nil {
		t.Fatalf("cloudimageeditorcore.MountImageByNbd : %s\n", err.Error())
	}
	defer cloudimageeditorcore.UnmountImageFreeNbd(nbdDev, mntRoot)

	t.Logf("TestHostname : nbdDev = %s, mntRoot = %s\n", nbdDev, mntRoot)

	err = cloudimageeditorcore.ExcuteModuleByName(mntRoot, "hostname", "TestHostname")
	if err != nil {
		t.Fatalf("cloudimageeditorcore.ExcuteModuleByName : %s\n", err.Error())
	}
}
