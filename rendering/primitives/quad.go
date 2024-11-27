package primitives

import "chemists-lab/rendering"

func GenQuad(l rendering.AttribLocator) rendering.Mesh {
	type Vec2 = rendering.Vec2
	verts := []vertex2{
		{pos: Vec2{0, 0}, uv: Vec2{0, 0}},
		{pos: Vec2{1, 0}, uv: Vec2{1, 0}},
		{pos: Vec2{1, 1}, uv: Vec2{1, 1}},
		{pos: Vec2{0, 1}, uv: Vec2{0, 1}},
	}

	indices := []uint8{
		0, 1, 2,
		0, 2, 3,
	}

	mesh := rendering.NewIndexedMesh(verts, indices, l)
	return &mesh
}
