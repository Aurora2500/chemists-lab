package rendering

import "github.com/go-gl/gl/v4.3-core/gl"

type Texture2D struct {
	id id
}

func NewTexture() Texture2D {
	var id id
	gl.GenTextures(1, &id)

	return Texture2D{id: id}
}
