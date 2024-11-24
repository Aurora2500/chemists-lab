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
	type Quat = rendering.Quat
	type compound struct {
		rotation Quat
		pos      Vec3
		compound int32
	}
	compounds := []compound{
		{pos: Vec3{-3, 0, 0}, compound: 0, rotation: Quat{W: 1}},
		{pos: Vec3{3, -.3, 0}, compound: 0, rotation: Quat{W: 1}},
		{pos: Vec3{0, 0, 3}, compound: 1, rotation: Quat{W: 1}},
		{pos: Vec3{4, 2, -3}, compound: 1, rotation: Quat{W: 1}},
		{pos: Vec3{-2, -2, 5}, compound: 1, rotation: Quat{W: 1}},
		{pos: Vec3{4, 1, 2}, compound: 1, rotation: Quat{W: 1}},
	}

	ssbo := rendering.NewSsbo[compound]()
	ssbo.Allocate(len(compounds), rendering.STREAM_DRAW)
	ssbo.Update(compounds)
	pt := game.NewPeriodicTable()
	cinfo := game.NewCompoundInfo()

	var timer win.Timer

	for window.Running() {
		window.Clear()
		dt := timer.Tick()

		ccam.SetVP(s)
		pt.Ssbo.BindShader(0)
		cinfo.BindShader(1)
		ssbo.BindShader(2)
		sphere.DrawInstanced(int32(len(compounds)))

		for i := range compounds {
			compounds[i].pos = compounds[i].pos.Add(
				Vec3{
					rand.Float32(), rand.Float32(), rand.Float32(),
				}.Add(Vec3{-.5, -.5, -.5}).Mul(2 * float32(dt)),
			)

			compounds[i].rotation = compounds[i].rotation.Mul(
				rendering.RotateAround(float32(dt*float64(i)), Vec3{float32(i), 1, 1}.Normalize()),
			)
		}
		ssbo.Update(compounds)

		window.Swap()
	}
}
