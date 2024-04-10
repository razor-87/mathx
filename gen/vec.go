package gen

func Vec(size int) []float64 {
	vec := make([]float64, size)
	for i := range vec {
		vec[i] = float64(i + 1)
	}
	return vec
}

func VecRand(size int) []float64 {
	vec := make([]float64, size)
	for i := 0; i < size; i++ {
		vec[i] = r.Float64() + 0.001
	}
	return vec
}
