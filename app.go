package main

import (
	"chemists-lab/game"
	"chemists-lab/rendering"
	"chemists-lab/rendering/primitives"
	"chemists-lab/resources"
	win "chemists-lab/windowing"
	"math/rand"

	"github.com/go-gl/gl/v2.1/gl"
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

	sphere := primitives.GenIcosphere(5, s)
	positions := []rendering.Vec4{
		{0, 0, 0},
		{20, 0, 0},
		{-20, 0, 0},
		{0, 0, 20},
		{0, 0, -20},
	}

	ssbo := rendering.NewSsbo[rendering.Vec4]()
	ssbo.Allocate(len(positions), gl.STREAM_DRAW)
	ssbo.Update(positions)

	for window.Running() {
		window.Clear()

		ccam.SetVP(s)
		ssbo.BindShader(0)
		sphere.DrawInstanced(int32(len(positions)))

		for i := range positions {
			positions[i][0] += (rand.Float32() - 0.5) * 0.2
			positions[i][1] += (rand.Float32() - 0.5) * 0.2
			positions[i][2] += (rand.Float32() - 0.5) * 0.2
		}
		ssbo.Update(positions)

		window.Swap()
	}
}
