package vec

import "github.com/razor-87/mathx"

func FillZeroes[T mathx.Vector[E], E mathx.Floaty](v T) {
	clear(v)
}

func FillOnes[T mathx.Vector[E], E mathx.Floaty](v T) {
	FillBy(v, 1)
}

func FillBy[T mathx.Vector[E], E mathx.Floaty](v T, value E) {
	for i := range v {
		v[i] = value
	}
}

func FillFn[T mathx.Vector[E], E mathx.Floaty](v T, f func(E) E) {
	for i := range v {
		v[i] = f(v[i])
	}
}
