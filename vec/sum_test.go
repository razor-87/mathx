package vec

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	const epsilon = 1e-6
	tests := []struct {
		name    string
		xs      []float64
		wantRet float64
	}{
		{
			"empty",
			[]float64{},
			0,
		},
		{
			"one",
			[]float64{1.1},
			1.1,
		},
		{
			"two",
			[]float64{1.1, 2.2},
			3.3,
		},
		{
			"five",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5},
			16.6,
		},
		{
			"seven",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7},
			30.9,
		},
		{
			"eight",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7, 8.8},
			39.7,
		},
		{
			"nine",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7, 8.8, 9.9},
			49.6,
		},
		{
			"ten",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10},
			59.7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Sum(tt.xs); !(math.Abs(tt.wantRet-gotRet) < epsilon) {
				t.Errorf("Sum() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
