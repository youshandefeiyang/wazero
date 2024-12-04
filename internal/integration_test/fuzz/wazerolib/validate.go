package main

import (
	"context"

	"github.com/youshandefeiyang/wazero"
	"github.com/youshandefeiyang/wazero/api"
	"github.com/youshandefeiyang/wazero/experimental"
)

// Ensure that validation and compilation do not panic!
func tryCompile(wasmBin []byte) {
	ctx := context.Background()
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigCompiler().
		WithCoreFeatures(api.CoreFeaturesV2|experimental.CoreFeaturesThreads))
	defer func() {
		if err := r.Close(context.Background()); err != nil {
			panic(err)
		}
	}()
	compiled, err := r.CompileModule(ctx, wasmBin)
	if err == nil {
		if err := compiled.Close(context.Background()); err != nil {
			panic(err)
		}
	}
}
