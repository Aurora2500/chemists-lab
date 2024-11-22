package primitives

import (
	"chemists-lab/rendering"
	"math"
)

type vertex3 struct {
	pos rendering.Vec3
}

func cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}
