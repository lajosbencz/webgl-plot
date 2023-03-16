package pkg

import (
	"github.com/seqsense/webgl-go"
)

type Point struct {
	X, Y float32
}

func (p *Point) ToList() []float32 {
	return []float32{p.X, p.Y}
}

type Line []Point

func (p *Line) ToList() []float32 {
	r := []float32{}
	for _, v := range *p {
		r = append(r, v.ToList()...)
	}
	return r
}

type Color struct {
	R, G, B float32
}

func (p *Color) ToList() []float32 {
	return []float32{p.R, p.G, p.B}
}

type LinePlot struct {
	values []float64
	// points  Line
	bgColor   Color
	lineColor Color
	gl        *webgl.WebGL
	vs        webgl.Shader
	fs        webgl.Shader
	program   webgl.Program
}

func NewLinePlot(gl *webgl.WebGL, bgColor Color) *LinePlot {
	return &LinePlot{
		gl:      gl,
		bgColor: bgColor,
		// points:  lineData,
		values: []float64{},
	}
}

func (r *LinePlot) ChangeBg(bg Color) {
	r.bgColor = bg
}

func (r *LinePlot) ChangeLine(c Color) {
	r.lineColor = c
}

func (r *LinePlot) FlowData(values []float64) {
	rn := len(r.values)
	n := len(values)
	if n > rn {
		n = rn
	}
	r.values = r.values[n:]
	r.values = append(r.values, values...)
}

func (r *LinePlot) SetData(values []float64) {
	r.values = values
}

func (r *LinePlot) AsPoints() Line {
	values := r.values
	n := len(values)
	if n < 1 {
		return Line{}
	}
	min := values[0]
	max := min
	points := Line{}
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	yd := max - min
	yr := 1. / float64(yd)
	xs := 1. / float64(n)
	for i, v := range values {
		if i == 0 {
			continue
		}
		v0 := float32(v*yr)*2 - 1
		v1 := float32(values[i-1]*yr)*2 - 1
		i0 := float32(xs*float64(i))*2 - 1
		i1 := float32(xs*float64(i-1))*2 - 1
		points = append(points, Point{X: i0, Y: v0})
		points = append(points, Point{X: i1, Y: v1})
	}
	return points
}

func (r *LinePlot) Init(vsSource, fsSource string) error {
	gl := r.gl

	// glParam := gl.GetParameter(ALIASED_LINE_WIDTH_RANGE)
	// lineWidthMin := glParam.Index(0).Float()
	// lineWidthMax := glParam.Index(1).Float()
	// log.Printf("line width min: %f max: %f\n", lineWidthMin, lineWidthMax)

	vs, err := InitVertexShader(gl, vsSource)
	if err != nil {
		return err
	}

	fs, err := InitFragmentShader(gl, fsSource)
	if err != nil {
		return err
	}

	program, err := LinkShaders(gl, nil, vs, fs)
	if err != nil {
		return err
	}
	gl.UseProgram(program)
	r.vs = vs
	r.fs = fs
	r.program = program
	return nil
}

func (r *LinePlot) Render() {
	gl := r.gl

	gl.ClearColor(r.bgColor.R, r.bgColor.G, r.bgColor.B, 0.9)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.DEPTH_TEST)

	width := gl.Canvas.ClientWidth()
	height := gl.Canvas.ClientHeight()
	gl.Viewport(0, 0, width, height)

	points := r.AsPoints()
	pointCount := len(points)
	if pointCount > 1 {

		// vertex
		vertexBuffer := gl.CreateBuffer()
		gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(points.ToList()), gl.STATIC_DRAW)

		position := gl.GetAttribLocation(r.program, "aPosition")
		gl.VertexAttribPointer(position, 2, gl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(position)

		// fragment
		colors := []float32{}
		for n := 0; n < pointCount; n++ {
			colors = append(colors, r.lineColor.ToList()...)
		}
		colorBuffer := gl.CreateBuffer()
		gl.BindBuffer(gl.ARRAY_BUFFER, colorBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(colors), gl.STATIC_DRAW)

		color := gl.GetAttribLocation(r.program, "aColor")
		gl.VertexAttribPointer(color, 3, gl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(color)
	}

	// render
	gl.DrawArrays(gl.LINES, 0, pointCount)
}
