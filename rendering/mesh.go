package rendering

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

type Mesh interface {
	Draw()
	DrawInstanced(n int32)
}

type ObjectMesh struct {
	vao *Vao
	vbo *Vbo
}

func NewObjectMesh[V any](verts []V, l AttribLocator) ObjectMesh {
	vbo, err := NewVbo(verts)
	if err != nil {
		panic(err)
	}
	vao, err := NewVao[V](l, vbo)
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

func NewIndexedMesh[V any, I OpenGLIndex](verts []V, indices []I, l AttribLocator) IndexedMesh {
	mesh := NewObjectMesh(verts, l)
	ebo := NewEBO(indices)
	mesh.vao.BindEbo(ebo)
	return IndexedMesh{
		ObjectMesh: mesh,
		ebo:        ebo,
	}
}

func (mesh *IndexedMesh) Draw() {
	mesh.vao.Bind()
	gl.DrawElementsWithOffset(gl.TRIANGLES, int32(mesh.ebo.length), mesh.ebo.indexType, 0)
}

func (mesh *IndexedMesh) DrawInstanced(n int32) {
	mesh.vao.Bind()
	gl.DrawElementsInstanced(gl.TRIANGLES, int32(mesh.ebo.length), mesh.ebo.indexType, gl.PtrOffset(0), n)
}

type InstancedMesh struct {
	IndexedMesh
}
