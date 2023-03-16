package main

import (
	"log"
	"syscall/js"
	"time"

	plot "github.com/lajosbencz/webgl-plot/pkg"
	webgl "github.com/seqsense/webgl-go"
)

var (
	bgColor = []float32{0.3, 0.3, 0.3}
)

const vsSource = `
attribute vec3 position;
attribute vec3 color;
varying vec3 vColor;
void main(void) {
  gl_Position = vec4(position, 1.0);
  vColor = color;
}
`

const fsSource = `
precision mediump float;
varying vec3 vColor;
void main(void) {
  gl_FragColor = vec4(vColor, 1.);
}
`

var vertices = []float32{
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0, 0.5, 0,
}

var colors = []float32{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

func changeBgColor(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return js.ValueOf(false)
	}
	bgColor[0] = float32(args[0].Float())
	bgColor[1] = float32(args[1].Float())
	bgColor[2] = float32(args[2].Float())
	return js.ValueOf(true)
}

func createTriangle(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 || args[0].IsUndefined() || args[0].IsNull() {
		return nil
	}
	canvas := args[0]
	gl, err := webgl.New(canvas)
	if err != nil {
		panic(err)
	}

	glParam := gl.GetParameter(plot.ALIASED_LINE_WIDTH_RANGE)
	lineWidthMin := glParam.Index(0).Float()
	lineWidthMax := glParam.Index(1).Float()
	log.Printf("line width min: %f max: %f\n", lineWidthMin, lineWidthMax)

	var vs, fs webgl.Shader
	if vs, err = plot.InitVertexShader(gl, vsSource); err != nil {
		panic(err)
	}

	if fs, err = plot.InitFragmentShader(gl, fsSource); err != nil {
		panic(err)
	}

	program, err := plot.LinkShaders(gl, nil, vs, fs)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	go func() {
		for {
			width := gl.Canvas.ClientWidth()
			height := gl.Canvas.ClientHeight()

			vertexBuffer := gl.CreateBuffer()
			gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
			gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(vertices), gl.STATIC_DRAW)

			colorBuffer := gl.CreateBuffer()
			gl.BindBuffer(gl.ARRAY_BUFFER, colorBuffer)
			gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(colors), gl.STATIC_DRAW)

			gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
			position := gl.GetAttribLocation(program, "position")
			gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
			gl.EnableVertexAttribArray(position)

			gl.BindBuffer(gl.ARRAY_BUFFER, colorBuffer)
			color := gl.GetAttribLocation(program, "color")
			gl.VertexAttribPointer(color, 3, gl.FLOAT, false, 0, 0)
			gl.EnableVertexAttribArray(color)

			gl.ClearColor(bgColor[0], bgColor[1], bgColor[2], 0.9)
			gl.Clear(gl.COLOR_BUFFER_BIT)
			gl.Enable(gl.DEPTH_TEST)
			gl.Viewport(0, 0, width, height)
			gl.DrawArrays(gl.TRIANGLES, 0, len(vertices)/3)

			time.Sleep(time.Millisecond * 500)
		}
	}()

	return js.ValueOf(map[string]interface{}{
		"changeBgColor": js.FuncOf(changeBgColor),
	})
}
