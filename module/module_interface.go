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

// getArgsFromConf : get interface, address, netmask, gateway from config map
func getArgsFromConf(ifaceConf map[string]string) (string, string, string, string, error) {
	iface, ok := ifaceConf["interface"]
	if !ok {
		return "", "", "", "", fmt.Errorf("can not get interface")
	}

	address, ok := ifaceConf["address"]
	if !ok {
		return "", "", "", "", fmt.Errorf("can not get address")
	}

	netmask, ok := ifaceConf["netmask"]
	if !ok {
		return "", "", "", "", fmt.Errorf("can not get netmask")
	}

	gateway, ok := ifaceConf["gateway"]
	if !ok {
		return "", "", "", "", fmt.Errorf("can not get gateway")
	}

	return iface, address, netmask, gateway, nil
}

// setUbuntuInterface : set interface for ubuntu. Ex, write config to /etc/network/interfaces.d/enp2s1
func setUbuntuInterface(root string, ifaceConf map[string]string) error {
	iface, address, netmask, gateway, err := getArgsFromConf(ifaceConf)
	if err != nil {
		return err
	}

	ifacePath := root + "/etc/network/interfaces.d/" + iface
	ifaceConfig := fmt.Sprintf(`
auto %s
iface %s inet static
address %s
netmask %s
gateway %s`, iface, iface, address, netmask, gateway)

	return ioutil.WriteFile(ifacePath, []byte(ifaceConfig), 0644)
}

// setCentosInterface : set interface for centos. Ex, write config to /etc/sysconfig/network-scripts/ifcfg-enp2s1
func setCentosInterface(root string, ifaceConf map[string]string) error {
	iface, address, netmask, gateway, err := getArgsFromConf(ifaceConf)
	if err != nil {
		return err
	}

	ifacePath := root + "/etc/sysconfig/network-scripts/ifcfg-" + iface
	ifaceConfig := fmt.Sprintf(`
NAME=%s
BOOTPROTO=static
ONBOOT=yes
IPADDR=%s
NETMASK=%s
GATEWAY=%s`, iface, address, netmask, gateway)

	return ioutil.WriteFile(ifacePath, []byte(ifaceConfig), 0644)
}

// InterfaceExecute : interface module execute method, config address, netmask, gateway...
func InterfaceExecute(env map[string]string, params interface{}) error {
	osType, ok := env["os"]
	if !ok {
		return fmt.Errorf("can not get os type")
	}

	root, ok := env["root"]
	if !ok {
		return fmt.Errorf("can not get mount root type")
	}

	ifaceConf := params.(map[string]string)

	switch osType {
	case cloudimageeditorcore.OSCentos:
		return setCentosInterface(root, ifaceConf)

	case cloudimageeditorcore.OSUbuntu:
		return setUbuntuInterface(root, ifaceConf)

	case cloudimageeditorcore.OSUnknown:
		fallthrough

	default:
		return fmt.Errorf("unsupported os type : %s", osType)
	}
}
