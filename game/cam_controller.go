package game

import (
	"chemists-lab/rendering"
	events "chemists-lab/windowing"
	"cmp"
	"math"
)

func clamp[T cmp.Ordered](x, lo, hi T) T {
	return max(min(x, hi), lo)
}

type CamController struct {
	Cam          rendering.OrbitCamera
	Lens         rendering.PerspectiveLens
	Sensitivity  float32
	dragging     bool
	xlast, ylast float32
}

func (cc *CamController) RegisterCallbacks(reg events.Callbacker) {
	reg.RegisterMouseButton(func(button events.MouseButton, action events.Action, mod events.ModifierKey) {
		if button == events.MouseButtonRight {
			cc.dragging = (action == events.Press)
		}
	})
	reg.RegisterMousePos(func(xpos, ypos float64) {
		if cc.dragging {
			dragx := cc.xlast - float32(xpos)
			dragy := cc.ylast - float32(ypos)
			cc.Cam.Yaw -= cc.Sensitivity * dragx
			cc.Cam.Yaw = float32(math.Mod(float64(cc.Cam.Yaw), 2*math.Pi))
			cc.Cam.Pitch -= cc.Sensitivity * dragy
			cc.Cam.Pitch = clamp(cc.Cam.Pitch, -math.Pi/2, math.Pi/2)
		}
		cc.xlast, cc.ylast = float32(xpos), float32(ypos)
	})
	reg.RegisterScroll(func(xoff, yoff float64) {
		cc.Cam.Distance = clamp(cc.Cam.Distance-float32(yoff), 1, 100)
	})
}

func (cc *CamController) SetVP(s *rendering.Shader) {
	s.SetUniformMat4("view", cc.Cam.View())
	s.SetUniformMat4("proj", cc.Lens.Projection())
}
