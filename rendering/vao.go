package rendering

import (
	"errors"
	"reflect"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const attrib_tag = "attrib"

var (
	ErrNonStructVertex = errors.New("rendering: non struct type used as a vertex")
	ErrUntaggedField   = errors.New("rendering: struct with untagged field used as vertex")
	ErrUnhandledType   = errors.New("rendering: type used as attrib is not handled")
)

type Vao struct {
	id id
}

func (vao *Vao) Id() id {
	return vao.id
}

func (vao *Vao) Bind() {
	gl.BindVertexArray(vao.id)
}

func (vao *Vao) Delete() {
	gl.DeleteVertexArrays(1, &vao.id)
}

func (vao *Vao) BindEbo(ebo *Ebo) {
	gl.VertexArrayElementBuffer(vao.id, ebo.id)
}

func set_attrib(loc id, vertex reflect.Type, field reflect.StructField) error {

	var num int32
	switch field.Type {
	case reflect.TypeOf((*float32)(nil)).Elem():
		num = 1
	case reflect.TypeOf((*mgl32.Vec2)(nil)).Elem():
		num = 2
	case reflect.TypeOf((*mgl32.Vec3)(nil)).Elem():
		num = 3
	case reflect.TypeOf((*mgl32.Vec4)(nil)).Elem():
		num = 4
	default:
		return ErrUnhandledType
	}
	gl.VertexAttribPointerWithOffset(loc, num, gl.FLOAT, false, int32(vertex.Size()), field.Offset)
	return nil
}

func NewVao[T any](program *Shader, vbo *Vbo) (*Vao, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Struct {
		return nil, ErrNonStructVertex
	}

	var vao Vao
	gl.GenVertexArrays(1, &vao.id)
	gl.BindVertexArray(vao.id)
	vbo.Bind()
	gl.VertexArrayVertexBuffer(vao.id, 0, vbo.id, 0, int32(t.Size()))

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get(attrib_tag)
		if name == "" {
			name = f.Name
		}
		loc := uint32(gl.GetAttribLocation(program.id, gl.Str(name+"\x00")))
		gl.VertexArrayAttribBinding(vao.id, loc, 0)
		gl.EnableVertexArrayAttrib(vao.id, loc)
		set_attrib(loc, t, f)
	}

	return &vao, nil
}
