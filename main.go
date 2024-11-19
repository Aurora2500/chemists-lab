package main

import (
	win "chemists-lab/windowing"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
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

	gl.ClearColor(0.3, 0.3, 0.7, 1.0)
	for window.Running() {
		window.Clear()

		window.Swap()
	}
}
