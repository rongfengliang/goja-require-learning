package main

import (
	"fmt"

	"github.com/dop251/goja"
)

// MyVM myjvm
type MyVM struct {
	jsRuntime *goja.Runtime
	script    string
}

func (vm *MyVM) init() {
	vm.jsRuntime.Set("login", func(name, password string) string {
		return fmt.Sprintf("%s-%s", name, password)
	})
	// load underscore for global user
	m, _ :=myrequire.Require("underscore-min.js")
	vm.jsRuntime.Set("_", m)
}

// Exec script
func (vm *MyVM) Exec() (goja.Value, error) {
	return vm.jsRuntime.RunString(vm.script)
}
