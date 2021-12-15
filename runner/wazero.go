package runner

import (
	"os"
	"testing"

	"github.com/tetratelabs/wazero/wasi"
	"github.com/tetratelabs/wazero/wasm"
	"github.com/tetratelabs/wazero/wasm/wazeroir"
)

// newWASMStoreWithWazero prepare for wazero wasm store.
func newWASMStoreWithWazero(t testing.TB, wasmFile string) *wasm.Store {
	binary, err := os.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}
	mod, err := wasm.DecodeModule(binary)
	if err != nil {
		t.Fatal(err)
	}

	store := wasm.NewStore(wazeroir.NewEngine())
	if err := wasi.NewEnvironment().Register(store); err != nil {
		t.Fatal(err)
	}

	if err := store.Instantiate(mod, wazeroModName); err != nil {
		t.Fatal(err)
	}

	return store
}

// callWASMFuncWithWazero call test func with wazero loader.
func callWASMFuncWithWazero(t testing.TB, store *wasm.Store, in int32) int32 {
	ret, retTypes, err := store.CallFunction(wazeroModName, fibFuncName, uint64(in))
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}

	if len(ret) != len(retTypes) {
		t.Fatalf("got values length %d is not equal with got value types length %d", len(ret), len(retTypes))
	}
	if retTypes[0] != wasm.ValueTypeI32 {
		t.Fatalf("got value types[0] is %v, but want %v", retTypes[0], wasm.ValueTypeI32)
	}

	return int32(ret[0])
}
