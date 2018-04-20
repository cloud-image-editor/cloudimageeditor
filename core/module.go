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

	cloudimageeditorutil "github.com/cloud-image-editor/cloudimageeditor/util"
)

// ExecuteFunc : module excute function type
type ExecuteFunc func(env map[string]string, params interface{}) error

// RevertFunc : module revert function type
type RevertFunc func(params interface{}) error

// Module : module base struct
type Module struct {
	Name    string
	Execute ExecuteFunc
	Revert  RevertFunc
}

// ModulesMap : modules map, module name -> mudule struct
var ModulesMap map[string]Module

// moduleMap : get ModulesMap, if ModulesMap is nil, make new map
func moduleMap() map[string]Module {
	if ModulesMap == nil {
		ModulesMap = make(map[string]Module)
	}

	return ModulesMap
}

// ModuleRegister : register module
func ModuleRegister(module Module) {
	_, ok := moduleMap()[module.Name]
	if ok {
		cloudimageeditorutil.LogD("module %s has already registered, update!\n", module.Name)
	}

	moduleMap()[module.Name] = module
}

// ExcuteModuleByName : excute module method. image mount point as @root, try to search module by @name
func ExcuteModuleByName(root string, name string, params interface{}) error {
	_, ok := moduleMap()[name]
	if !ok {
		return fmt.Errorf("module %s is missing", name)
	}

	env := map[string]string{
		"root": root,
		"os":   OSGuess(root),
	}
	err := moduleMap()[name].Execute(env, params)

	return err
}
