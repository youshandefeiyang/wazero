package expctxkeys

// FunctionListenerFactoryKey is a context.Context Value key.
// Its associated value should be a FunctionListenerFactory.
//
// See https://github.com/youshandefeiyang/wazero/issues/451
type FunctionListenerFactoryKey struct{}
