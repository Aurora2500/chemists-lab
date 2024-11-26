package rendering

import (
	"image"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type Texture2D struct {
	id id
}

func NewTexture() Texture2D {
	var id id
	gl.GenTextures(1, &id)

	return Texture2D{id: id}
}

func (t *Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func (t *Texture2D) BindUnit(unit id) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func (t *Texture2D) Upload(img *image.Alpha) {
	width, height := int32(img.Stride), int32(len(img.Pix))/int32(img.Stride)
	t.Bind()
	gl.TexStorage2D(gl.TEXTURE_2D, 1, gl.R8, width, height)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, width, height, gl.RED, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
}
