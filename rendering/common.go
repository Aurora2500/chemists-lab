package rendering

import "github.com/go-gl/mathgl/mgl32"

type id = uint32

type (
	Vec2 = mgl32.Vec2
	Vec3 = mgl32.Vec3
	Vec4 = mgl32.Vec4
	Mat4 = mgl32.Mat4
)

type Resource interface {
	Delete()
}
