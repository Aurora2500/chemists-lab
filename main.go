package main

import (
	"chemists-lab/rendering"
	win "chemists-lab/windowing"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window, err := win.CreateWindow("Chemist's Lab")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	var s rendering.Shader

	type Vert2D struct {
		pos mgl32.Vec2 `attrib:"position"`
		uv  mgl32.Vec2
	}
	_, err = rendering.CreateVAO[Vert2D](&s)
	if err != nil {
		panic(err)
	}

	gl.ClearColor(0.3, 0.3, 0.7, 1.0)
	for window.Running() {
		window.Clear()

		var foo Vert2D
		foo.pos = mgl32.Vec2{3}
		foo.uv = mgl32.Vec2{3}

		window.Swap()
	}
}
