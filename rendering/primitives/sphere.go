package primitives

import (
	"chemists-lab/rendering"
	"math"
)

func midpoint(a, b rendering.Vec3) rendering.Vec3 {
	return a.Add(b).Mul(0.5).Normalize()
}

func GenIcosphere(subdivs int, s *rendering.Shader) rendering.Mesh {
	phi := float32(math.Phi)

	vertices := []rendering.Vec3{
		{-1, phi, 0}, {1, phi, 0}, {-1, -phi, 0}, {1, -phi, 0},
		{0, -1, phi}, {0, 1, phi}, {0, -1, -phi}, {0, 1, -phi},
		{phi, 0, -1}, {phi, 0, 1}, {-phi, 0, -1}, {-phi, 0, 1},
	}

	faces := [][3]uint16{
		{0, 11, 5}, {0, 5, 1}, {0, 1, 7}, {0, 7, 10}, {0, 10, 11},
		{1, 5, 9}, {5, 11, 4}, {11, 10, 2}, {10, 7, 6}, {7, 1, 8},
		{3, 9, 4}, {3, 4, 2}, {3, 2, 6}, {3, 6, 8}, {3, 8, 9},
		{4, 9, 5}, {2, 4, 11}, {6, 2, 10}, {8, 6, 7}, {9, 8, 1},
	}
	for i := range vertices {
		vertices[i] = vertices[i].Normalize()
	}

	for i := 0; i < subdivs; i++ {
		newFaces := [][3]uint16{}
		midpoints := make(map[[2]uint16]uint16)

		genMidpoint := func(v1, v2 uint16) uint16 {
			edge := [2]uint16{v1, v2}
			if v1 > v2 {
				edge = [2]uint16{v2, v1}
			}

			if idx, ok := midpoints[edge]; ok {
				return idx
			}

			mid := midpoint(vertices[v1], vertices[v2])
			vertices = append(vertices, mid)
			idx := uint16(len(vertices) - 1)
			midpoints[edge] = idx
			return idx
		}

		for _, face := range faces {
			v0, v1, v2 := face[0], face[1], face[2]
			a := genMidpoint(v0, v1)
			b := genMidpoint(v1, v2)
			c := genMidpoint(v0, v2)

			newFaces = append(newFaces,
				[3]uint16{v0, a, c},
				[3]uint16{v1, b, a},
				[3]uint16{v2, c, b},
				[3]uint16{a, b, c},
			)
		}
		faces = newFaces
	}

	verts := make([]vertex3, len(vertices))
	for i := range vertices {
		verts[i].pos = vertices[i]
	}
	indices := make([]uint16, 3*len(faces))
	for i := range faces {
		indices[3*i+0] = faces[i][0]
		indices[3*i+1] = faces[i][1]
		indices[3*i+2] = faces[i][2]
	}

	mesh := rendering.NewIndexedMesh(verts, indices, s)
	return &mesh
}
