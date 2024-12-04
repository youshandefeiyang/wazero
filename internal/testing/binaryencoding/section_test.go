package binaryencoding

import (
	"testing"

	"github.com/youshandefeiyang/wazero/internal/testing/require"
	"github.com/youshandefeiyang/wazero/internal/wasm"
)

func TestEncodeFunctionSection(t *testing.T) {
	require.Equal(t, []byte{wasm.SectionIDFunction, 0x2, 0x01, 0x05}, EncodeFunctionSection([]wasm.Index{5}))
}

// TestEncodeStartSection uses the same index as TestEncodeFunctionSection to highlight the encoding is different.
func TestEncodeStartSection(t *testing.T) {
	require.Equal(t, []byte{wasm.SectionIDStart, 0x01, 0x05}, EncodeStartSection(5))
}
