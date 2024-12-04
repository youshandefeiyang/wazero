package v2

import (
	"context"
	"testing"

	"github.com/youshandefeiyang/wazero"
	"github.com/youshandefeiyang/wazero/api"
	"github.com/youshandefeiyang/wazero/internal/integration_test/spectest"
	"github.com/youshandefeiyang/wazero/internal/platform"
)

const enabledFeatures = api.CoreFeaturesV2

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}
