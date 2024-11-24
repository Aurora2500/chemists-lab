package rendering

import "github.com/go-gl/gl/v4.3-core/gl"

type Ssbo struct {
	id id
}

func NewSsbo() *Ssbo {
	return &Ssbo{}
}

func (ssbo *Ssbo) Delete() {
	gl.DeleteBuffers(1, &ssbo.id)
}
