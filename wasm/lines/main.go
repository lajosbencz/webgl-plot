package main

import "syscall/js"

var (
	jsGlobal js.Value
)

func setup() {
	jsGlobal = js.Global()
	jsGlobal.Set("createLinePlot", js.FuncOf(createLinePlot))
}

func run() {
}

func main() {
	setup()
	go run()
	select {}
}
