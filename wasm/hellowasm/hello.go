package main

import "syscall/js"

func hellowasm(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	str := "Hello from Go! You said: " + message
	jsGlobal.Get("console").Call("log", str)
	return str
}

func hellovalue(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(map[string]interface{}{"foo": js.ValueOf([]interface{}{"bar", "baz", "bax"})})
}
