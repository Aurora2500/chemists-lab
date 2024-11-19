package windowing

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	window *glfw.Window
}

func CreateWindow(title string) (*Window, error) {
	glfw.Init()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(1600, 900, title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return nil, err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version)

	return &Window{window: window}, nil
}

func (win *Window) Destroy() {

}

func (win *Window) Running() bool {
	return !win.window.ShouldClose()
}

func (win *Window) Stop() {
	win.window.SetShouldClose(true)
}

func (win *Window) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (win *Window) Swap() {
	if win.window.GetKey(glfw.KeyQ) == glfw.Press {
		win.Stop()
	}
	win.window.SwapBuffers()
	glfw.PollEvents()
}
