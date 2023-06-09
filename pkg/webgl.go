package pkg

// Taken with minor modifications from https://github.com/seqsense/pcdeditor/blob/master/shader_js.go

import (
	"errors"
	"fmt"
	"syscall/js"

	webgl "github.com/seqsense/webgl-go"
)

const (
	ALIASED_LINE_WIDTH_RANGE = 0x846E
	ALIASED_POINT_SIZE_RANGE = 0x846D
)

func NewCanvas(this js.Value, args []js.Value) (*webgl.WebGL, error) {
	if len(args) < 1 || args[0].IsUndefined() || args[0].IsNull() {
		return nil, fmt.Errorf("missing first argument: element")
	}
	canvas := args[0]
	gl, err := webgl.New(canvas)
	if err != nil {
		return nil, err
	}
	return gl, nil
}

func InitVertexShader(gl *webgl.WebGL, src string) (webgl.Shader, error) {
	s := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(s, src)
	gl.CompileShader(s)
	if !gl.GetShaderParameter(s, gl.COMPILE_STATUS).(bool) {
		compilationLog := gl.GetShaderInfoLog(s)
		return webgl.Shader(js.Null()), fmt.Errorf("compile failed (VERTEX_SHADER) %v", compilationLog)
	}
	return s, nil
}

func InitFragmentShader(gl *webgl.WebGL, src string) (webgl.Shader, error) {
	s := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(s, src)
	gl.CompileShader(s)
	if !gl.GetShaderParameter(s, gl.COMPILE_STATUS).(bool) {
		compilationLog := gl.GetShaderInfoLog(s)
		return webgl.Shader(js.Null()), fmt.Errorf("compile failed (FRAGMENT_SHADER) %v", compilationLog)
	}
	return s, nil
}

func LinkShaders(gl *webgl.WebGL, fbVarings []string, shaders ...webgl.Shader) (webgl.Program, error) {
	program := gl.CreateProgram()
	for _, s := range shaders {
		gl.AttachShader(program, s)
	}
	if len(fbVarings) > 0 {
		gl.TransformFeedbackVaryings(program, fbVarings, gl.SEPARATE_ATTRIBS)
	}
	gl.LinkProgram(program)
	if !gl.GetProgramParameter(program, gl.LINK_STATUS).(bool) {
		return webgl.Program(js.Null()), errors.New("link failed: " + gl.GetProgramInfoLog(program))
	}
	return program, nil
}
