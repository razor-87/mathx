package vec

import (
	"math"

	"github.com/razor-87/mathx"
)

const stride = 8

func Add[T mathx.Vector[E], E mathx.Floaty](v, w T) {
	for ; len(v) >= stride && len(w) >= stride; v, w = v[stride:], w[stride:] {
		v[0] += w[0]
		v[1] += w[1]
		v[2] += w[2]
		v[3] += w[3]
		v[4] += w[4]
		v[5] += w[5]
		v[6] += w[6]
		v[7] += w[7]
	}
	for i := 0; i < len(v) && i < len(w); i++ {
		v[i] += w[i]
	}
}

func AddScaled[T mathx.Vector[E], E mathx.Floaty](v, w T, c E) {
	for ; len(v) >= stride && len(w) >= stride; v, w = v[stride:], w[stride:] {
		v[0] += w[0] * c
		v[1] += w[1] * c
		v[2] += w[2] * c
		v[3] += w[3] * c
		v[4] += w[4] * c
		v[5] += w[5] * c
		v[6] += w[6] * c
		v[7] += w[7] * c
	}
	for i := 0; i < len(v) && i < len(w); i++ {
		v[i] += w[i] * c
	}
}

func DotProd[T mathx.Vector[E], E mathx.Floaty](v, w T) (ret E) {
	for ; len(v) >= stride && len(w) >= stride; v, w = v[stride:], w[stride:] {
		ret += v[0]*w[0] + v[1]*w[1] + v[2]*w[2] + v[3]*w[3] + v[4]*w[4] + v[5]*w[5] + v[6]*w[6] + v[7]*w[7]
	}
	for i := 0; i < len(v) && i < len(w); i++ {
		ret += v[i] * w[i]
	}

	return ret
}

func Length[T mathx.Vector[E], E mathx.Floaty](v T) (ret E) {
	return E(math.Pow(float64(DotProd(v, v)), 0.5))
}

func Scale[T mathx.Vector[E], E mathx.Floaty](v T, c E) {
	for ; len(v) >= stride; v = v[stride:] {
		v[0] *= c
		v[1] *= c
		v[2] *= c
		v[3] *= c
		v[4] *= c
		v[5] *= c
		v[6] *= c
		v[7] *= c
	}
	for i := range v {
		v[i] *= c
	}
}

func Sum[T mathx.Vector[E], E mathx.Floaty](v T) (ret E) {
	for ; len(v) >= stride; v = v[stride:] {
		ret += v[0] + v[1] + v[2] + v[3] + v[4] + v[5] + v[6] + v[7]
	}
	for i := range v {
		ret += v[i]
	}

	return ret
}

func Unit[T mathx.Vector[E], E mathx.Floaty](v T) {
	Scale(v, 1/Length(v))
}
