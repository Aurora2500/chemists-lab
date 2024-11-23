package main

import (
	"chemists-lab/rendering"
	"chemists-lab/rendering/primitives"
	"chemists-lab/resources"
	win "chemists-lab/windowing"
	"cmp"
	"math"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func clamp[T cmp.Ordered](x, lo, hi T) T {
	return max(min(x, hi), lo)
}

const sensitivity = 0.01

func runApp() {
	window, err := win.CreateWindow("Chemist's Lab")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	manager := resources.NewManager("assets")

	cam := rendering.OrbitCamera{}
	cam.Distance = 30
	var drag bool = false
	var dragPos rendering.Vec2

	width, height := window.Size()
	lens := rendering.PerspectiveLens{
		Near:   0.1,
		Far:    500,
		Width:  uint32(width),
		Height: uint32(height),
		Fov:    60.,
	}

	window.MouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		if button == glfw.MouseButtonRight {
			drag = (action == glfw.Press)
		}
	})
	window.MouseCallback(func(w *glfw.Window, xpos, ypos float64) {
		if drag {
			dragx := dragPos.X() - float32(xpos)
			dragy := dragPos.Y() - float32(ypos)
			cam.Yaw -= float32(sensitivity * dragx)
			cam.Yaw = float32(math.Mod(float64(cam.Yaw), 2*math.Pi))
			cam.Pitch -= float32(sensitivity * dragy)
			cam.Pitch = clamp(cam.Pitch, -math.Pi/2, math.Pi/2)
		}
		dragPos = rendering.Vec2{float32(xpos), float32(ypos)}
	})
	window.MouseScrollCallback(func(w *glfw.Window, xoff, yoff float64) {
		cam.Distance = clamp(cam.Distance-float32(yoff), 1, 100)
	})

	s := manager.GetShader("sphere")
	s.Use()

	sphere := primitives.GenIcosphere(3, s)

	positions := []rendering.Vec3{
		{0, 0, 0},
		{2, 2, 1},
		{20, 0, 0},
		{-20, 0, 0},
		{0, 0, 20},
		{0, 0, -20},
	}

	for window.Running() {
		window.Clear()

		s.SetUniformMat4("view", cam.View())
		s.SetUniformMat4("proj", lens.Projection())
		for _, pos := range positions {
			model := mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())
			s.SetUniformMat4("model", model)
			sphere.Draw()
		}

		window.Swap()
	}
}
