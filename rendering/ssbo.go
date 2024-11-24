package rendering

import (
	"reflect"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Ssbo[T any] struct {
	id id
}

func NewSsbo[T any]() *Ssbo[T] {
	var id id
	gl.GenBuffers(1, &id)
	return &Ssbo[T]{id: id}
}

func (ssbo *Ssbo[T]) Bind() {
	gl.BindBuffer(gl.SHADER_STORAGE_BUFFER, ssbo.id)
}

func (ssbo *Ssbo[T]) Allocate(size int, use uint32) {
	ssbo.Bind()
	size_bytes := size * int(reflect.TypeFor[T]().Size())
	gl.BufferData(gl.SHADER_STORAGE_BUFFER, size_bytes, nil, use)
}

func (ssbo *Ssbo[T]) Update(data []T) {
	ssbo.Bind()
	size_bytes := len(data) * int(reflect.TypeFor[T]().Size())
	gl.BufferSubData(gl.SHADER_STORAGE_BUFFER, 0, size_bytes, gl.Ptr(data))
}

func (ssbo *Ssbo[T]) BindShader(n uint32) {
	gl.BindBufferBase(gl.SHADER_STORAGE_BUFFER, n, ssbo.id)
}

func (ssbo *Ssbo[T]) Delete() {
	gl.DeleteBuffers(1, &ssbo.id)
}
