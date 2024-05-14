package vec

import "github.com/razor-87/mathx"

const stride = 8

func Add[T mathx.Vector[E], E mathx.Floaty](v, w T) {
	for i := 0; i < len(v) && i < len(w); i++ {
		v[i] += w[i]
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

func Scale[T mathx.Vector[E], E mathx.Floaty](v T, c E) {
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
