package gen

import "mathx/types"

func Vec[T types.Vector[E], E types.Floaty](size int) T {
	vec := make(T, size)
	for i := range vec {
		vec[i] = E(i + 1)
	}
	return vec
}

func VecRand[T types.Vector[E], E types.Floaty](size int) T {
	vec := make(T, size)
	for i := 0; i < size; i++ {
		vec[i] = E(r.Float64() + 0.001)
	}
	return vec
}
