package windowing

import "github.com/go-gl/glfw/v3.2/glfw"

type MouseButton = glfw.MouseButton

const (
	MouseButtonLeft  MouseButton = glfw.MouseButtonLeft
	MouseButtonRight             = glfw.MouseButtonRight
)

type Action = glfw.Action

const (
	Press   Action = glfw.Press
	Release        = glfw.Release
)

type ModifierKey = glfw.ModifierKey
