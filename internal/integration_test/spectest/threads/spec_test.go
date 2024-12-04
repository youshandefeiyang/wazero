package spectest

import (
	"context"
	"embed"
	"testing"

	"github.com/youshandefeiyang/wazero"
	"github.com/youshandefeiyang/wazero/api"
	"github.com/youshandefeiyang/wazero/experimental"
	"github.com/youshandefeiyang/wazero/internal/integration_test/spectest"
	"github.com/youshandefeiyang/wazero/internal/platform"
)

//go:embed testdata/*.wasm
//go:embed testdata/*.json
var testcases embed.FS

const enabledFeatures = api.CoreFeaturesV2 | experimental.CoreFeaturesThreads

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}
