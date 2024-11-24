package rendering

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type id = uint32

type (
	Vec2 = mgl32.Vec2
	Vec3 = mgl32.Vec3
	Vec4 = mgl32.Vec4
	Mat4 = mgl32.Mat4
)

const (
	STATIC_DRAW  = gl.STATIC_DRAW
	DYNAMIC_DRAW = gl.DYNAMIC_DRAW
	STREAM_DRAW  = gl.STREAM_DRAW
)

type Resource interface {
	Delete()
}
