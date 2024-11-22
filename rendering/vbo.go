package rendering

import (
	"reflect"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Vbo struct {
	id id
}

func (vbo *Vbo) Id() id {
	return vbo.id
}

func (vbo *Vbo) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
}

func NewVbo[T any](data []T) (*Vbo, error) {
	v := reflect.TypeOf(data).Elem()

	var vbo Vbo

	gl.GenBuffers(1, &vbo.id)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
	gl.BufferData(gl.ARRAY_BUFFER, int(v.Size())*len(data), gl.Ptr(data), gl.STATIC_DRAW)

	return &vbo, nil
}
