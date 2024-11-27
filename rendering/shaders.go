package rendering

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Shader struct {
	id   id
	locs map[string]int32
}

type shaderStage struct {
	src   string
	stage id
}

func NewShader(src string) (*Shader, error) {
	var stages []shaderStage

	for _, line := range strings.Split(src, "\n") {
		if strings.HasPrefix(line, "//shader") {
			var stageId id
			switch line[len("//shader "):] {
			case "vertex":
				stageId = gl.VERTEX_SHADER
			case "geometry":
				stageId = gl.GEOMETRY_SHADER
			case "fragment":
				stageId = gl.FRAGMENT_SHADER
			}
			stages = append(stages, shaderStage{stage: stageId})
		} else {
			if len(stages) == 0 {
				continue
			}
			stage := &stages[len(stages)-1]
			stage.src += line + "\n"
		}
	}
	var shader Shader
	compiledStages := make([]id, len(stages))
	for i := range stages {
		var err error
		compiledStages[i], err = compileShader(stages[i].src, stages[i].stage)
		if err != nil {
			return nil, err
		}
	}
	shader.id = gl.CreateProgram()

	for _, stage := range compiledStages {
		gl.AttachShader(shader.id, stage)
	}
	gl.LinkProgram(shader.id)

	var status int32
	gl.GetProgramiv(shader.id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shader.id, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shader.id, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	for _, stage := range compiledStages {
		gl.DeleteShader(stage)
	}

	shader.locs = make(map[string]int32)
	return &shader, nil
}

func compileShader(src string, shaderType id) (id, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(src + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile %v: %v", src, log)
	}

	return shader, nil
}

func (s *Shader) Id() id {
	return s.id
}

func (s *Shader) Use() {
	gl.UseProgram(s.id)
}

func (s *Shader) Delete() {
	gl.DeleteProgram(s.id)
}

func (s *Shader) get_uniform_loc(name string) int32 {
	loc, ok := s.locs[name]
	if !ok {
		loc = gl.GetUniformLocation(s.id, gl.Str(name+"\x00"))
		s.locs[name] = loc
	}
	return loc
}

func (s *Shader) SetUniformMat4(uniform string, x Mat4) {
	loc := s.get_uniform_loc(uniform)
	gl.ProgramUniformMatrix4fv(s.id, loc, 1, false, &x[0])
}

func (s *Shader) SetUniformTex2D(uniform string, tex *Texture2D, unit uint32) {
	loc := s.get_uniform_loc(uniform)
	tex.BindUnit(unit)
	gl.ProgramUniform1i(s.id, loc, int32(unit))
}

type AttribLocator interface {
	Locate(field reflect.StructField, idx int) int32
}

type ShaderLocator struct {
	Shader *Shader
}

func (sl ShaderLocator) Locate(field reflect.StructField, idx int) int32 {
	attrib := field.Tag.Get("attrib")
	if attrib == "" {
		attrib = field.Name
	}
	return gl.GetAttribLocation(sl.Shader.id, gl.Str(attrib+"\x00"))
}

type MapLocator map[string]int32

func (ml MapLocator) Locate(field reflect.StructField, idx int) int32 {
	attrib := field.Tag.Get("attrib")
	if attrib == "" {
		attrib = field.Name
	}
	loc, ok := ml[attrib]
	if !ok {
		return -1
	}
	return loc
}

type PosLocator struct{}

func (pl PosLocator) Locate(field reflect.StructField, idx int) int32 {
	return int32(idx)
}
