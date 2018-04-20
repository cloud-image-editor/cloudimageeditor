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
	cloudimageeditorutil "github.com/cloud-image-editor/cloudimageeditor/util"
)

// copyFolder : copy full folder(include symbolic link) by "cp -rd src dest"
func copyFolder(src string, dest string) error {
	dir, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, fi := range dir {
		result, err := cloudimageeditorutil.Execute("cp", []string{"-rd", src + "/" + fi.Name(), dest})
		if err != nil {
			return fmt.Errorf(err.Error() + " : " + result)
		}
	}

	return nil
}

// FolderExecute : folder module execute method, usually copy src to from
func FolderExecute(env map[string]string, params interface{}) error {
	osType, ok := env["os"]
	if !ok {
		return fmt.Errorf("can not get os type")
	}

	root, ok := env["root"]
	if !ok {
		return fmt.Errorf("can not get mount root type")
	}

	parasMap := params.(map[string]string)
	src, ok := parasMap["src"]
	if !ok {
		return fmt.Errorf("can not get src folder")
	}

	dest, ok := parasMap["dest"]
	if !ok {
		return fmt.Errorf("can not get dest folder")
	}

	switch osType {
	case cloudimageeditorcore.OSCentos:
		fallthrough

	case cloudimageeditorcore.OSUbuntu:
		return copyFolder(src, root+"/"+dest)

	case cloudimageeditorcore.OSUnknown:
		fallthrough

	default:
		return fmt.Errorf("unsupported os type : %s", osType)
	}
}
