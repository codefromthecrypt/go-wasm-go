package jsgoja

import (
	_ "embed"

	"github.com/dop251/goja"
)

//go:embed goja.js
var jsScript string

func newGojaVm() *goja.Runtime {
	vm := goja.New()

	if _, err := vm.RunString(jsScript); err != nil {
		panic(err)
	}

	return vm
}

func exportFn(fnName string, to interface{}) {
	vm := newGojaVm()

	err := vm.ExportTo(vm.Get(fnName), to)
	if err != nil {
		panic(err)
	}
}

func NewFibonacci() func(int32) int32 {
	var fn func(int32) int32
	exportFn("fibonacci", &fn)
	return fn
}

func NewHTTPBasicAuth() func(string, string) {
	var fn func(string, string)
	exportFn("httpbasicAuth", &fn)
	return fn
}
