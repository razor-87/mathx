package vec

func Sum(xs []float64) (ret float64) {
	for i := range xs {
		ret += xs[i]
	}
	return ret
}
