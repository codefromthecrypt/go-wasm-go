package main

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/provider/jsgoja"
)

func Benchmark_fibonacci_single_10(b *testing.B) {
	benchmark_fibonacci_single(b, 1)
}

func Benchmark_fibonacci_single_20(b *testing.B) {
	benchmark_fibonacci_single(b, 20)
}

func Benchmark_fibonacci_single_30(b *testing.B) {
	benchmark_fibonacci_single(b, 30)
}

func Benchmark_fibonacci_single_40(b *testing.B) {
	benchmark_fibonacci_single(b, 40)
}
func Benchmark_fibonacci_paralle_10(b *testing.B) {
	benchmark_fibonacci_paralle(b, 10)
}

func Benchmark_fibonacci_paralle_20(b *testing.B) {
	benchmark_fibonacci_paralle(b, 20)
}

func Benchmark_fibonacci_paralle_30(b *testing.B) {
	benchmark_fibonacci_paralle(b, 30)
}

func Benchmark_fibonacci_paralle_40(b *testing.B) {
	benchmark_fibonacci_paralle(b, 40)
}

func benchmark_fibonacci_paralle(b *testing.B, fbIn uint32) {
	// go gc 会导致 wasmer 异常,需要手动线停止GC.
	debug.SetGCPercent(-1)
	b.ResetTimer()

	b.Run(fmt.Sprintf("native - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fibonacci(fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("plugin - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			soFn := newGoPluginFunc(b, goPluginSo, wasmFuncName)
			for pb.Next() {
				soFn(fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("wasm-wazero - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			// 必须在线程里面加载, 不能在线程外加载，然后并发.
			store := newWASMStoreWithWazero(b, wasmTinygo)
			for pb.Next() {
				callWASMFuncWithWazero(b, store, fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("wasm-wasmer - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			wasmFn := getWasmFuncWithWasmer(b, wasmTinygo)
			for pb.Next() {
				callWASMFuncWithWasmer(b, wasmFn, fbIn)
			}
		})
	})
	b.Run(fmt.Sprintf("wasm-wasmedge - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			vm, conf := getWasmedgeInstance(b, wasmTinygo)
			defer vm.Release()
			defer conf.Release()
			for pb.Next() {
				callWASMFuncWithWasmedge(b, vm, fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("js - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			f := jsgoja.NewFibonacci()
			for pb.Next() {
				f(fbIn)
			}
		})
	})
}

func benchmark_fibonacci_single(b *testing.B, fbIn uint32) {
	b.Run(fmt.Sprintf("native - fb(%d)", fbIn), func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			fibonacci(fbIn)
		}
	})

	b.Run(fmt.Sprintf("plugin - fb(%d)", fbIn), func(b *testing.B) {
		soFn := newGoPluginFunc(b, goPluginSo, wasmFuncName)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			soFn(fbIn)
		}
	})

	b.Run(fmt.Sprintf("wasm-wazero - fb(%d)", fbIn), func(b *testing.B) {
		store := newWASMStoreWithWazero(b, wasmTinygo)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			callWASMFuncWithWazero(b, store, fbIn)
		}
	})

	b.Run(fmt.Sprintf("wasm-wasmer - fb(%d)", fbIn), func(b *testing.B) {
		wasmFn := getWasmFuncWithWasmer(b, wasmTinygo)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			callWASMFuncWithWasmer(b, wasmFn, fbIn)
		}
	})
	b.Run(fmt.Sprintf("wasm-wasmedge - fb(%d)", fbIn), func(b *testing.B) {
		vm, conf := getWasmedgeInstance(b, wasmTinygo)
		defer vm.Release()
		defer conf.Release()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			callWASMFuncWithWasmedge(b, vm, fbIn)
		}
	})

	b.Run(fmt.Sprintf("js - fb(%d)", fbIn), func(b *testing.B) {
		f := jsgoja.NewFibonacci()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			f(fbIn)
		}
	})
}
