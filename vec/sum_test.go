package vec

import (
	"math"
	"mathx/gen"
	"testing"
)

func TestSum(t *testing.T) {
	const epsilon = 1e-6
	tests := []struct {
		name    string
		xs      []float64
		wantRet float64
		eps     bool
	}{
		{
			name: "empty",
		},
		{
			"one",
			[]float64{1.1},
			1.1,
			true,
		},
		{
			"two",
			[]float64{1.1, 2.2},
			3.3,
			true,
		},
		{
			"five",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5},
			16.6,
			true,
		},
		{
			"seven",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7},
			30.9,
			true,
		},
		{
			"eight",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7, 8.8},
			39.7,
			true,
		},
		{
			"nine",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7, 8.8, 9.9},
			49.6,
			true,
		},
		{
			"ten",
			[]float64{1.1, 2.2, 3.3, 4.5, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10},
			59.7,
			true,
		},
		{
			"thirty one",
			gen.Vec(31),
			496,
			false,
		},
		{
			"thirty two",
			gen.Vec(32),
			528,
			false,
		},
		{
			"thirty three",
			gen.Vec(33),
			561,
			false,
		},
		{
			"ninety nine",
			gen.Vec(99),
			4950,
			false,
		},
		{
			"hundred",
			gen.Vec(100),
			5050,
			false,
		},
		{
			"hundred and one",
			gen.Vec(101),
			5151,
			false,
		},
		{
			"thousand",
			gen.Vec(1000),
			500500,
			false,
		},
		{
			"ten thousand",
			gen.Vec(10000),
			50005000,
			false,
		},
	}
	var (
		gotRet float64
		equal  bool
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet = Sum(tt.xs)
			if equal = func() bool {
				if tt.eps {
					return math.Abs(tt.wantRet-gotRet) < epsilon
				}
				return gotRet == tt.wantRet
			}(); !equal {
				t.Errorf("Sum() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
