package main

import "syscall/js"

var (
	jsGlobal js.Value
)

func setup() {
	jsGlobal = js.Global()
	jsGlobal.Set("hellowasm", js.FuncOf(hellowasm))
	jsGlobal.Set("hellovalue", js.FuncOf(hellovalue))
}

func run() {
}

func main() {
	setup()
	go run()
	select {}
}
