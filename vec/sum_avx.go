package vec

func SumAVX(xs []float64) float64 {
	const k = 64
	var rem float64
	if r := len(xs) % k; r != 0 {
		rem = Sum(xs[:r])
		xs = xs[r:]
	}
	if len(xs) == 0 {
		return rem
	}

	return rem + sumAVXFloat64(xs)
}

func sumAVXFloat64(x []float64) float64
