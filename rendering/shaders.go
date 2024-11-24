package rendering

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Shader struct {
	id   id
	locs map[string]int32
}

func NewShader(src string) (*Shader, error) {
	var vertSrc, fragSrc string
	var currSrc *string

	for _, line := range strings.Split(src, "\n") {
		if line == "//shader vertex" {
			currSrc = &vertSrc
			continue
		} else if line == "//shader fragment" {
			currSrc = &fragSrc
			continue
		}
		if currSrc == nil {
			continue
		}

		*currSrc = *currSrc + line + "\n"
	}

	var shader Shader
	vert, err := compileShader(vertSrc, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	frag, err := compileShader(fragSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	shader.id = gl.CreateProgram()

	gl.AttachShader(shader.id, vert)
	gl.AttachShader(shader.id, frag)
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

	gl.DeleteShader(vert)
	gl.DeleteShader(frag)

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
		gpu_src := gl.GoStr(*csources)
		return 0, fmt.Errorf("failed to compile %v: %v", gpu_src, log)
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
	loc := int32(s.get_uniform_loc(uniform))
	gl.ProgramUniformMatrix4fv(s.id, loc, 1, false, &x[0])
}
