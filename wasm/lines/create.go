package main

import (
	"fmt"
	"syscall/js"

	"github.com/lajosbencz/webgl-plot/pkg"
	plot "github.com/lajosbencz/webgl-plot/pkg"
)

func createLinePlot(this js.Value, args []js.Value) any {
	gl, err := plot.NewCanvas(this, args)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	bgColor := pkg.Color{R: 0.1, G: 0.1, B: 0.1}
	lp := pkg.NewLinePlot(gl, bgColor)
	lp.Init(vsSourceColor, fsSourceColor)
	lp.Render()

	// go func(lp *pkg.LinePlot) {
	// 	for {
	// 		lp.Render()
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }(lp)

	changeBgColor := js.FuncOf(func(this js.Value, args []js.Value) any {
		lp.ChangeBg(plot.Color{
			R: float32(args[0].Float()),
			G: float32(args[1].Float()),
			B: float32(args[2].Float()),
		})
		return nil
	})

	changeLineColor := js.FuncOf(func(this js.Value, args []js.Value) any {
		lp.ChangeLine(plot.Color{
			R: float32(args[0].Float()),
			G: float32(args[1].Float()),
			B: float32(args[2].Float()),
		})
		return nil
	})

	render := js.FuncOf(func(this js.Value, args []js.Value) any {
		lp.Render()
		return nil
	})

	setData := js.FuncOf(func(this js.Value, args []js.Value) any {
		rawData := []float64{}
		for _, av := range args {
			v := av.Float()
			rawData = append(rawData, v)
		}
		lp.SetData(rawData)
		return nil
	})

	flowData := js.FuncOf(func(this js.Value, args []js.Value) any {
		rawData := []float64{}
		for _, av := range args {
			v := av.Float()
			rawData = append(rawData, v)
		}
		lp.FlowData(rawData)
		return nil
	})

	jsObj := js.ValueOf(map[string]any{
		"changeBgColor":   changeBgColor,
		"changeLineColor": changeLineColor,
		"setData":         setData,
		"flowData":        flowData,
		"render":          render,
	})

	return jsObj
}
