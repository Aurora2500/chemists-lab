package rendering

import "github.com/go-gl/gl/v4.1-core/gl"

type Shader struct {
	id   id
	locs map[string]id
}

func NewShader(vertSrc, fragSrc string) (*Shader, error) {
	return nil, nil
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

func (s *Shader) get_attrib_loc(name string) id {
	loc, ok := s.locs[name]
	if !ok {
		loc = uint32(gl.GetAttribLocation(s.id, gl.Str(name+"\x00")))
	}
	return loc
}
