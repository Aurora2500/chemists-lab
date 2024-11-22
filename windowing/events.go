package windowing

import "github.com/go-gl/glfw/v3.2/glfw"

func (win *Window) MouseCallback(cb glfw.CursorPosCallback) {
	win.window.SetCursorPosCallback(cb)
}

func (win *Window) MouseButtonCallback(cb glfw.MouseButtonCallback) {
	win.window.SetMouseButtonCallback(cb)
}
