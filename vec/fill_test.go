package vec

import "testing"

func TestFillZeroes(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Rand[[]float64](5)
		FillZeroes(v)
		for i := range v {
			if v[i] != 0 {
				t.Error("expected 0")
			}
		}
	})
}

func TestFillOnes(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Rand[[]float64](5)
		FillOnes(v)
		for i := range v {
			if v[i] != 1 {
				t.Error("expected 1")
			}
		}
	})
}

func TestFillBy(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Rand[[]float64](5)
		FillBy(v, 2)
		for i := range v {
			if v[i] != 2 {
				t.Error("expected 2")
			}
		}
	})
}

func TestFillFn(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Inc[[]float64](5)
		FillFn(v, func(x float64) float64 { return x + 1 })
		for i := range v {
			if n := float64(i + 2); v[i] != n {
				t.Errorf("expected %f, got %f", n, v[i])
			}
		}
	})
}
