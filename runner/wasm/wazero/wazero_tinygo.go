package wazero

import (
	"context"
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi"
)

const wazeroModName = "wasmtest"

// NewWASMStoreWithWazero prepare for wazero wasm runtime.
func NewWASMStoreWithWazero(b testing.TB, wasmFile string) (api.Module, func() error) {
	ctx := context.Background()

	source, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	runtime := wazero.NewRuntime()

	wm, err := wasi.InstantiateSnapshotPreview1(ctx, runtime)
	if err != nil {
		_ = wm.Close(ctx)
		b.Fatal(err)
	}

	compiled, err := runtime.CompileModule(ctx, source, wazero.NewCompileConfig())
	if err != nil {
		_ = wm.Close(ctx)
		b.Fatal(err)
	}

	config := wazero.NewModuleConfig().WithName(wazeroModName)
	module, err := runtime.InstantiateModule(ctx, compiled, config)
	if err != nil {
		_ = wm.Close(ctx)
		b.Fatal(err)
	}

	return module, func() (err error) {
		if e := module.Close(ctx); e != nil {
			err = e
		}
		if e := wm.Close(ctx); e != nil && err != nil {
			err = e
		}
		return
	}
}

// CallWASMFuncWithWazero call test func with wazero loader.
func CallWASMFuncWithWazero(t testing.TB, module api.Module, funcName string, args ...uint64) []uint64 {
	ret, err := module.ExportedFunction(funcName).Call(context.Background(), args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}
