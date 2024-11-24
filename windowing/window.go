package windowing

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	window *glfw.Window
	cbr    *CallbackRegistry
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
		// fmt.Printf("Source: 0x%x\n", source)
		// fmt.Printf("Type: 0x%x\n", gltype)
		// fmt.Printf("ID: %d\n", id)
		// fmt.Printf("Severity: 0x%x\n", severity)
		fmt.Printf("OpenGL: %s\n", message)
	}, nil)
	gl.DebugMessageControl(gl.DEBUG_SOURCE_API, gl.DEBUG_TYPE_OTHER, gl.DONT_CARE, 0, nil, false)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	var win *Window = &Window{window: window, cbr: &CallbackRegistry{}}

	gl.ClearColor(0.8, 0.8, 0.8, 1.0)
	var geometry_limit int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_OUTPUT_VERTICES, &geometry_limit)
	var geometry_instance_limit int32
	gl.GetIntegerv(gl.MAX_GEOMETRY_SHADER_INVOCATIONS, &geometry_instance_limit)
	println("max geometry output vertices", geometry_limit)
	println("max geometry instancing", geometry_instance_limit)

	win.setupRegistry()

	return win, nil
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
