package rendering

import (
	"errors"
	"reflect"

	"github.com/go-gl/gl/v4.1-core/gl"
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

func CreateVAO[T any](program *Shader) (*Vao, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Struct {
		return nil, ErrNonStructVertex
	}

	var vao Vao
	gl.GenVertexArrays(1, &vao.id)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get(attrib_tag)
		if name == "" {
			name = f.Name
		}
		println("field name:", name)
		loc := uint32(gl.GetAttribLocation(program.id, gl.Str(name+"\x00")))
		gl.EnableVertexArrayAttrib(vao.id, loc)
		set_attrib(loc, t, f)
	}

	return &vao, nil
}
