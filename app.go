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
	compounds := []game.Compound{
		{Pos: Vec3{-3, 0, 0}, Compound: 0, Rotation: Quat{W: 1}},
		{Pos: Vec3{3, -.3, 0}, Compound: 0, Rotation: Quat{W: 1}},
		{Pos: Vec3{0, 0, 3}, Compound: 1, Rotation: Quat{W: 1}},
		{Pos: Vec3{4, 2, -3}, Compound: 1, Rotation: Quat{W: 1}},
		{Pos: Vec3{-2, -2, 5}, Compound: 1, Rotation: Quat{W: 1}},
		{Pos: Vec3{4, 1, 2}, Compound: 1, Rotation: Quat{W: 1}},
	}

	system := game.NewSystem(compounds)

	var timer win.Timer

	for window.Running() {
		window.Clear()
		dt := timer.Tick()

		ccam.SetVP(s)
		system.Bind()
		sphere.DrawInstanced(int32(len(compounds)))
		var i int = 0
		system.Compounds.Update(func(c *game.Compound) {
			c.Pos = c.Pos.Add(
				Vec3{
					rand.Float32(), rand.Float32(), rand.Float32(),
				}.Add(Vec3{-.5, -.5, -.5}).Mul(2 * float32(dt)),
			)
			c.Rotation = c.Rotation.Mul(
				rendering.RotateAround(
					10*float32(dt)*rand.Float32(),
					Vec3{rand.Float32(), rand.Float32(), rand.Float32()}.Add(Vec3{-.5, -.5, -.5}).Normalize(),
				),
			)
			i += 1
		})

		window.Swap()
	}
}
