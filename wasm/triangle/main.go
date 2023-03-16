package main

import (
	"syscall/js"
)

var (
	jsGlobal js.Value
)

func setup() {
	jsGlobal = js.Global()
	jsGlobal.Set("createTriangle", js.FuncOf(createTriangle))
}

func main() {
	go setup()
	select {}
}
