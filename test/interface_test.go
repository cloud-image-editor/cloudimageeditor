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

func TestInterface(t *testing.T) {
	cloudimageeditormodule.MuduleRegisterAll()

	imagePath := "ubuntu.qcow2"
	nbdDev, mntRoot, err := cloudimageeditorcore.MountImageByNbd(imagePath)
	if err != nil {
		t.Fatalf("cloudimageeditorcore.MountImageByNbd : %s\n", err.Error())
	}
	defer cloudimageeditorcore.UnmountImageFreeNbd(nbdDev, mntRoot)

	fmt.Printf("TestInterface : nbdDev = %s, mntRoot = %s\n", nbdDev, mntRoot)

	ifaceConf := make(map[string]string)
	ifaceConf["interface"] = "enp2s1"
	ifaceConf["address"] = "172.168.100.100"
	ifaceConf["netmask"] = "255.255.255.0"
	ifaceConf["gateway"] = "172.168.100.1"

	err = cloudimageeditorcore.ExcuteModuleByName(mntRoot, "interface", ifaceConf)
	if err != nil {
		t.Fatalf("cloudimageeditorcore.ExcuteModuleByName : %s\n", err.Error())
	}
}
