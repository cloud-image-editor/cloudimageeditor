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
	cloudimageeditorcore "github.com/cloud-image-editor/cloudimageeditor/core"
)

// MuduleRegisterAll : register all modules
func MuduleRegisterAll() {
	moduleFolder := cloudimageeditorcore.Module{
		Name:    "folder",
		Execute: FolderExecute,
	}
	cloudimageeditorcore.ModuleRegister(moduleFolder)

	moduleHostname := cloudimageeditorcore.Module{
		Name:    "hostname",
		Execute: HostnameExecute,
	}
	cloudimageeditorcore.ModuleRegister(moduleHostname)

	moduleInterface := cloudimageeditorcore.Module{
		Name:    "interface",
		Execute: InterfaceExecute,
	}
	cloudimageeditorcore.ModuleRegister(moduleInterface)
}
