package frontend

import (
	"github.com/youshandefeiyang/wazero/internal/engine/wazevo/ssa"
	"github.com/youshandefeiyang/wazero/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
