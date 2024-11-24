package main

import (
	"chemists-lab/game"
	"chemists-lab/rendering"
	"chemists-lab/rendering/primitives"
	"chemists-lab/resources"
	win "chemists-lab/windowing"
	"math/rand"
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
	type Vec3 = rendering.Vec3
	type atom struct {
		pos          Vec3
		atomicNumber int32
	}
	atoms := []atom{
		{pos: Vec3{0, 0, 0}, atomicNumber: 0},
		{pos: Vec3{2, 0, 0}, atomicNumber: 1},
		{pos: Vec3{-2, 0, 0}, atomicNumber: 0},
		{pos: Vec3{0, 0, 2}, atomicNumber: 0},
		{pos: Vec3{0, 0, -2}, atomicNumber: 1},
	}

	ssbo := rendering.NewSsbo[atom]()
	ssbo.Allocate(len(atoms), rendering.STREAM_DRAW)
	ssbo.Update(atoms)
	pt := game.NewPeriodicTable()

	var timer win.Timer

	for window.Running() {
		window.Clear()
		dt := timer.Tick()

		ccam.SetVP(s)
		pt.Ssbo.BindShader(0)
		ssbo.BindShader(1)
		sphere.DrawInstanced(int32(len(atoms)))

		for i := range atoms {
			atoms[i].pos[0] += (rand.Float32() - 0.5) * 2 * float32(dt)
			atoms[i].pos[1] += (rand.Float32() - 0.5) * 2 * float32(dt)
			atoms[i].pos[2] += (rand.Float32() - 0.5) * 2 * float32(dt)
		}
		ssbo.Update(atoms)

		window.Swap()
	}
}
