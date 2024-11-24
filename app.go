package main

import (
	"chemists-lab/game"
	"chemists-lab/rendering"
	"chemists-lab/rendering/primitives"
	"chemists-lab/resources"
	win "chemists-lab/windowing"

	"github.com/go-gl/mathgl/mgl32"
)

func runApp() {
	window, err := win.CreateWindow("Chemist's Lab")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	manager := resources.NewManager("assets")

	cam := rendering.OrbitCamera{}
	cam.Distance = 10
	width, height := window.Size()
	lens := rendering.PerspectiveLens{
		Near:   0.1,
		Far:    500,
		Width:  uint32(width),
		Height: uint32(height),
		Fov:    60.,
	}

	ccam := game.CamController{
		Cam:         cam,
		Lens:        lens,
		Sensitivity: 0.01,
	}
	ccam.RegisterCallbacks(window.CallbackRegistry())

	s := manager.GetShader("sphere")
	s.Use()

	sphere := primitives.GenIcosphere(3, s)

	positions := []rendering.Vec3{
		{0, 0, 0},
		{20, 0, 0},
		{-20, 0, 0},
		{0, 0, 20},
		{0, 0, -20},
	}

	for window.Running() {
		window.Clear()

		ccam.SetVP(s)
		for _, pos := range positions {
			model := mgl32.Translate3D(pos.X(), pos.Y(), pos.Z())
			s.SetUniformMat4("model", model)
			sphere.Draw()
		}

		window.Swap()
	}
}
