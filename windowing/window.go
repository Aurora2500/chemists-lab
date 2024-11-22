package windowing

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	window *glfw.Window
}

func CreateWindow(title string) (*Window, error) {
	glfw.Init()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
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

	gl.Enable(gl.DEBUG_OUTPUT)
	gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
	gl.DebugMessageCallback(func(
		source uint32,
		gltype uint32,
		id uint32,
		severity uint32,
		length int32,
		message string,
		userParam unsafe.Pointer) {
		fmt.Printf("OpenGL Debug Message\n")
		// fmt.Printf("Source: 0x%x\n", source)
		// fmt.Printf("Type: 0x%x\n", gltype)
		// fmt.Printf("ID: %d\n", id)
		// fmt.Printf("Severity: 0x%x\n", severity)
		fmt.Printf("Message: %s\n", message)
	}, nil)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
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
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (win *Window) Swap() {
	if win.window.GetKey(glfw.KeyQ) == glfw.Press {
		win.Stop()
	}
	win.window.SwapBuffers()
	glfw.PollEvents()
}

func (win *Window) Size() (int, int) {
	return win.window.GetSize()
}
