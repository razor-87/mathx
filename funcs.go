package mathx

import (
	"cmp"
	"math"
)

func IsClose[T Floaty](x, y T) bool {
	const epsilon = 1e-6
	if cmp.Compare(x, y) == 0 {
		return true
	}
	return math.Abs(float64(x-y)) < epsilon
}
