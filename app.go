package main

import (
	"chemists-lab/rendering"
	"chemists-lab/rendering/primitives"
	"chemists-lab/resources"
	win "chemists-lab/windowing"
	"cmp"
	"math"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func clamp[T cmp.Ordered](x, lo, hi T) T {
	return max(min(x, hi), lo)
}

func runApp() {
	window, err := win.CreateWindow("Chemist's Lab")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	manager := resources.NewManager("assets")

	cam := rendering.OrbitCamera{}
	cam.Distance = -20
	var drag bool = false

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
			cam.Yaw += float32(xpos)
			cam.Pitch = clamp(cam.Pitch+float32(ypos), -math.Pi/2, math.Pi/2)
		}
	})

	s := manager.GetShader("sphere")
	s.Use()

	sphere := primitives.GenIcosphere(3, s)

	for window.Running() {
		window.Clear()

		vp := lens.Projection().Mul4(cam.View())
		s.SetUniformMat4("vp", vp)

		sphere.Draw()

		window.Swap()
	}
}
