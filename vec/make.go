// Package vec provides functions for creating and manipulating vector data structures.
package vec

import (
	"math/rand"

	"github.com/razor-87/mathx"
)

func Zeros[T mathx.Vector[E], E mathx.Floaty](size int) (v T) {
	v = make(T, size)
	return v
}

func Ones[T mathx.Vector[E], E mathx.Floaty](size int) (v T) {
	v = Zeros[T](size)
	FillOnes(v)
	return v
}

func Rep[T mathx.Vector[E], E mathx.Floaty](size int, value E) (v T) {
	v = Zeros[T](size)
	FillBy(v, value)
	return v
}

func Inc[T mathx.Vector[E], E mathx.Floaty](size int) (v T) {
	v = Zeros[T](size)
	for i := range v {
		v[i] = E(i + 1)
	}
	return v
}

func IncBy[T mathx.Vector[E], E mathx.Floaty](size int, value E) (v T) {
	v = Zeros[T](size)
	for i, acc := 0, value; i < len(v); i, acc = i+1, acc+value {
		v[i] = acc
	}
	return v
}

func Rand[T mathx.Vector[E], E mathx.Floaty](size int) (v T) {
	const n = 1 << 53
	r := rand.New(rand.NewSource(rand.Int63())) //nolint:gosec
	v = Zeros[T](size)
	for i := range v {
		v[i] = E(r.Int63()&(n-1)) / n
	}

	return v
}

func Copy[T mathx.Vector[E], E mathx.Floaty](v T) (w T) {
	w = Zeros[T](len(v))
	_ = copy(w, v)
	return w
}
