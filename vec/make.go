package vec

import (
	"math/rand"

	"github.com/razor-87/mathx/types"
)

func Zeros[T types.Vector[E], E types.Floaty](size int) T {
	return make(T, size)
}

func Ones[T types.Vector[E], E types.Floaty](size int) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = E(1)
	}
	return vec
}

func Rep[T types.Vector[E], E types.Floaty](size int, value E) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = value
	}
	return vec
}

func Inc[T types.Vector[E], E types.Floaty](size int) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = E(i + 1)
	}
	return vec
}

func IncBy[T types.Vector[E], E types.Floaty](size int, value E) T {
	vec := Inc[T](size)
	for i := range vec {
		vec[i] *= value
	}
	return vec
}

func Rand[T types.Vector[E], E types.Floaty](size int) T {
	vec := Zeros[T](size)
	for i := range vec {
		vec[i] = E(rand.Float64()) //nolint:gosec
	}
	return vec
}
