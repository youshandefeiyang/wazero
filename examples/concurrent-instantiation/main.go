package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/youshandefeiyang/wazero"
)

// addWasm was generated by the following:
//
//	wasm-tools parse testdata/add.wat -o testdata/add.wasm
//
//go:embed testdata/add.wasm
var addWasm []byte

func main() {
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Compile the Wasm binary once so that we can skip the entire compilation time during instantiation.
	compiledWasm, err := r.CompileModule(ctx, addWasm)
	if err != nil {
		log.Panicf("failed to compile Wasm binary: %v", err)
	}

	var wg sync.WaitGroup
	const goroutines = 50
	wg.Add(goroutines)

	// Instantiate the Wasm module from `compiledWsam`, and invoke the exported "add" function concurrently.
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()

			// Instantiate a new Wasm module from the already compiled `compiledWasm`.
			instance, err := r.InstantiateModule(ctx, compiledWasm, wazero.NewModuleConfig().WithName(""))
			if err != nil {
				log.Panicf("[%d] failed to instantiate %v", i, err)
			}
			defer instance.Close(ctx) // This closes everything this Instance created.

			// Calculates "i + i" by invoking the exported "add" function.
			result, err := instance.ExportedFunction("add").Call(ctx, uint64(i), uint64(i))
			if err != nil {
				log.Panicf("[%d] failed to invoke \"add\": %v", i, err)
			}

			// Ensure the addition "i + i" is actually calculated.
			expected := uint64(i * 2)
			if result[0] != expected {
				log.Panicf("expected %d, but got %d", expected, result[0])
			}

			// Logs the result.
			fmt.Println(expected)
		}(i)
	}

	wg.Wait()
}
