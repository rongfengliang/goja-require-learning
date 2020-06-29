package main

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

var (
	myvm      MyVM
	myrequire *require.RequireModule
)

func init() {
	// registry := new(require.Registry)
	// use custom srcloader
	registry := require.NewRegistryWithLoader(func(path string) ([]byte, error) {
		return Asset(path)
	})
	myvm = MyVM{
		jsRuntime: goja.New(),
		script:    `require("dalong.js")("dalong","ddddd")`,
	}
	myrequire = registry.Enable(myvm.jsRuntime)
	myvm.init()
}
