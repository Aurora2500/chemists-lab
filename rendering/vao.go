package rendering

import (
	"errors"
	"reflect"

	"github.com/go-gl/gl/v4.3-core/gl"
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
	case reflect.TypeFor[float32]():
		num = 1
	case reflect.TypeFor[Vec2]():
		num = 2
	case reflect.TypeFor[Vec3]():
		num = 3
	case reflect.TypeFor[Vec4]():
		num = 4
	default:
		return ErrUnhandledType
	}
	gl.VertexAttribPointerWithOffset(loc, num, gl.FLOAT, false, int32(vertex.Size()), field.Offset)
	return nil
}

func NewVao[T any](locator AttribLocator, vbo *Vbo) (*Vao, error) {
	t := reflect.TypeFor[T]()
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
		floc := locator.Locate(f, i)
		if floc < 0 {
			panic("attribute location not found")
		}
		loc := uint32(floc)
		gl.VertexArrayAttribBinding(vao.id, loc, 0)
		gl.EnableVertexArrayAttrib(vao.id, loc)
		set_attrib(loc, t, f)
	}

	return &vao, nil
}
