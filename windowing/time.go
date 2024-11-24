package windowing

import "github.com/go-gl/glfw/v3.2/glfw"

type Timer struct {
	time float64
}

func (t *Timer) Tick() float64 {
	currTime := glfw.GetTime()
	dt := currTime - t.time
	t.time = currTime
	return dt
}
