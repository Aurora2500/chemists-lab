package rendering

import "github.com/go-gl/mathgl/mgl32"

type Camera interface {
	View() mgl32.Mat4
}

type lens interface {
	Projection() mgl32.Mat4
}

type OrbitCamera struct {
	Focus    mgl32.Vec3
	Yaw      float32
	Pitch    float32
	Distance float32
}

func (cam *OrbitCamera) View() mgl32.Mat4 {
	trans := mgl32.Translate3D(0, 0, -cam.Distance)
	trans = mgl32.HomogRotate3DX(cam.Pitch).Mul4(trans)
	trans = mgl32.HomogRotate3DY(cam.Yaw).Mul4(trans)
	trans = mgl32.Translate3D(cam.Focus.X(), cam.Focus.Y(), cam.Focus.Z()).Mul4(trans)
	return trans
}

type PerspectiveLens struct {
}
