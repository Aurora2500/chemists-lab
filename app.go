package main

import (
	"chemists-lab/rendering"
	win "chemists-lab/windowing"
	"os"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
)

type Vert struct {
	pos rendering.Vec2
	col rendering.Vec3
}

var trongle = []Vert{
	{pos: rendering.Vec2{-.5, -.5}, col: rendering.Vec3{1., 0., 0.}},
	{pos: rendering.Vec2{.5, -.5}, col: rendering.Vec3{0., 1., 0.}},
	{pos: rendering.Vec2{.0, .5}, col: rendering.Vec3{0., 0., 1.}},
}

func runApp() {
	window, err := win.CreateWindow("Chemist's Lab")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	tronglesrc, err := os.ReadFile("shaders/trongle.glsl")
	if err != nil {
		panic(err)
	}

	s, err := rendering.NewShader(string(tronglesrc))
	if err != nil {
		panic(err)
	}
	s.Use()
	vbo, err := rendering.NewVbo(trongle)
	if err != nil {
		panic(err)
	}
	vbo.Bind()
	vao, err := rendering.NewVao[Vert](s)
	if err != nil {
		panic(err)
	}
	gl.VertexArrayVertexBuffer(vao.Id(), 0, vbo.Id(), 0, 20)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	for window.Running() {
		window.Clear()

		vao.Bind()
		vbo.Bind()

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.Swap()
	}
}
