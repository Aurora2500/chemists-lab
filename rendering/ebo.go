package rendering

import (
	"reflect"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Ebo struct {
	id        id
	length    uint32
	indexType int
}

func NewEBO[I OpenGLIndex](indices []I) *Ebo {
	var id id
	gl.GenBuffers(1, &id)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, id)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(reflect.TypeFor[I]().Size()), gl.Ptr(indices), gl.STATIC_DRAW)

	return &Ebo{
		id:        id,
		length:    uint32(len(indices)),
		indexType: IndexType[I](),
	}
}

type OpenGLIndex interface {
	uint8 | uint16 | uint32
}

func IndexType[I OpenGLIndex]() int {
	t := reflect.TypeFor[I]()

	switch t {
	case reflect.TypeFor[uint8]():
		return gl.UNSIGNED_BYTE
	case reflect.TypeFor[uint16]():
		return gl.UNSIGNED_SHORT
	case reflect.TypeFor[uint32]():
		return gl.UNSIGNED_INT
	default:
		panic("unhandled index type")
	}
}
