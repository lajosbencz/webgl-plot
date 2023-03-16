package main

import "syscall/js"

var (
	jsGlobal js.Value
)

func setup() {
	jsGlobal = js.Global()
	jsGlobal.Set("goFunc", js.FuncOf(goFunc))
	jsGlobal.Set("createPlot", js.FuncOf(createPlot))
}

func goFunc(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	return "Hello from Go! You said: " + message
}

func createPlot(this js.Value, args []js.Value) interface{} {
	jsCanvas := args[0]
	obj := jsGlobal.Get("Object").New()
	for _, tag := range []string{"tagName", "id", "style"} {
		v := jsCanvas.Get(tag)
		if !v.IsUndefined() {
			obj.Set(tag, v.String())
		}
	}
	return obj
}

func run() {

}

func main() {
	setup()
	go run()
	select {}
}
