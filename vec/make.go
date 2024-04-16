// Package vec provides functions for creating and manipulating vector data structures.
package vec

import (
	"math/rand"

	"github.com/razor-87/mathx"
)

func Zeros[T mathx.Vector[E], E mathx.Floaty](size int) T {
	return make(T, size)
}

func Ones[T mathx.Vector[E], E mathx.Floaty](size int) T {
	return Rep[T](size, 1)
}

func Rep[T mathx.Vector[E], E mathx.Floaty](size int, value E) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = value
	}
	return vec
}

func Inc[T mathx.Vector[E], E mathx.Floaty](size int) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = E(i + 1)
	}
	return vec
}

func IncBy[T mathx.Vector[E], E mathx.Floaty](size int, value E) T {
	vec := Inc[T](size)
	for i := range vec {
		vec[i] *= value
	}
	return vec
}

func Rand[T mathx.Vector[E], E mathx.Floaty](size int) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = E(rand.Float64()) //nolint:gosec
	}
	return vec
}
