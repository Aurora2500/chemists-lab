package rendering

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

type Mesh interface {
	Draw()
}

type ObjectMesh struct {
	vao *Vao
	vbo *Vbo
}

func NewObjectMesh[V any](verts []V, s *Shader) ObjectMesh {
	vbo, err := NewVbo(verts)
	if err != nil {
		panic(err)
	}
	vao, err := NewVao[V](s, vbo)
	if err != nil {
		panic(err)
	}
	return ObjectMesh{
		vao: vao,
		vbo: vbo,
	}
}

type IndexedMesh struct {
	ObjectMesh
	ebo *Ebo
}

func NewIndexedMesh[V any, I OpenGLIndex](verts []V, indices []I, s *Shader) IndexedMesh {
	mesh := NewObjectMesh(verts, s)
	ebo := NewEBO(indices)
	mesh.vao.BindEbo(ebo)
	return IndexedMesh{
		ObjectMesh: mesh,
		ebo:        ebo,
	}
}

func (mesh *IndexedMesh) Draw() {
	mesh.vao.Bind()
	gl.DrawElementsWithOffset(gl.TRIANGLES, int32(mesh.ebo.length), uint32(mesh.ebo.indexType), 0)
}

type InstancedMesh struct {
	IndexedMesh
}
