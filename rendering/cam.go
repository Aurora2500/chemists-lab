package rendering

import "github.com/go-gl/mathgl/mgl32"

type Camera interface {
	View() Mat4
}

type Lens interface {
	Projection() Mat4
}

type OrbitCamera struct {
	Focus    Vec3
	Yaw      float32
	Pitch    float32
	Distance float32
}

func (cam *OrbitCamera) View() mgl32.Mat4 {
	return transPipeline(
		mgl32.Translate3D(0, 0, -cam.Distance),
		mgl32.HomogRotate3DX(cam.Pitch),
		mgl32.HomogRotate3DY(cam.Yaw),
		mgl32.Translate3D(-cam.Focus.X(), -cam.Focus.Y(), -cam.Focus.Z()),
	)
}

type PerspectiveLens struct {
	Near, Far     float32
	Width, Height uint32
	Fov           float32
}

func (lens *PerspectiveLens) Projection() Mat4 {
	aspect := float32(lens.Width) / float32(lens.Height)
	return mgl32.Perspective(mgl32.DegToRad(lens.Fov), aspect, lens.Near, lens.Far)
}

func transPipeline(transformations ...Mat4) Mat4 {
	mat := mgl32.Ident4()
	for _, t := range transformations {
		mat = mat.Mul4(t)
	}
	return mat
}
