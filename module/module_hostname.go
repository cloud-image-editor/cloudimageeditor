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
	"io/ioutil"

	cloudimageeditorcore "github.com/cloud-image-editor/cloudimageeditor/core"
)

func setHostname(filename string, hostname string) error {
	return ioutil.WriteFile(filename, []byte(hostname), 0644)
}

// HostnameExecute : hostname module execute method
func HostnameExecute(env map[string]string, params interface{}) error {
	osType, ok := env["os"]
	if !ok {
		return fmt.Errorf("can not get os type")
	}

	root, ok := env["root"]
	if !ok {
		return fmt.Errorf("can not get mount root type")
	}

	filename := root + "/etc/hostname"
	hostname := params.(string)

	switch osType {
	case cloudimageeditorcore.OSCentos:
		fallthrough

	case cloudimageeditorcore.OSUbuntu:
		return setHostname(filename, hostname)

	case cloudimageeditorcore.OSUnknown:
		fallthrough

	default:
		return fmt.Errorf("unsupported os type : %s", osType)
	}
}
