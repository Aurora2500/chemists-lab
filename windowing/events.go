package windowing

import "github.com/go-gl/glfw/v3.2/glfw"

type MousePosCallback = func(xpos, ypos float64)
type MouseButtonCallback = func(button glfw.MouseButton, action Action, mod ModifierKey)
type ScrollCallback = func(xoff, yoff float64)

type Callbacker interface {
	RegisterMousePos(cb MousePosCallback)
	RegisterMouseButton(cb MouseButtonCallback)
	RegisterScroll(cb ScrollCallback)
}

func (win *Window) setupRegistry() {
	win.window.SetCursorPosCallback(win.cbr.MousePosCallback)
	win.window.SetMouseButtonCallback(win.cbr.MouseButtonCallback)
	win.window.SetScrollCallback(win.cbr.ScrollCallback)
}

func (win *Window) CallbackRegistry() Callbacker {
	return win.cbr
}

type CallbackRegistry struct {
	mousePosCallbacks    []MousePosCallback
	mouseButtonCallbacks []MouseButtonCallback
	scrollCallbacks      []ScrollCallback
}

func (cbr *CallbackRegistry) RegisterMousePos(cb MousePosCallback) {
	cbr.mousePosCallbacks = append(cbr.mousePosCallbacks, cb)
}

func (cbr *CallbackRegistry) MousePosCallback(w *glfw.Window, xpos, ypos float64) {
	for _, cb := range cbr.mousePosCallbacks {
		cb(xpos, ypos)
	}
}

func (cbr *CallbackRegistry) RegisterMouseButton(cb MouseButtonCallback) {
	cbr.mouseButtonCallbacks = append(cbr.mouseButtonCallbacks, cb)
}

func (cbr *CallbackRegistry) MouseButtonCallback(w *glfw.Window, button MouseButton, action Action, mod ModifierKey) {
	for _, cb := range cbr.mouseButtonCallbacks {
		cb(button, action, mod)
	}
}

func (cbr *CallbackRegistry) RegisterScroll(cb ScrollCallback) {
	cbr.scrollCallbacks = append(cbr.scrollCallbacks, cb)
}

func (cbr *CallbackRegistry) ScrollCallback(w *glfw.Window, xoff, yoff float64) {
	for _, cb := range cbr.scrollCallbacks {
		cb(xoff, yoff)
	}
}
