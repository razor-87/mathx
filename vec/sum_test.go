package vec

import (
	"testing"

	"github.com/razor-87/mathx"
)

func TestSum(t *testing.T) {
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
			IncBy[[]float64](5, 1.1),
			16.5,
			true,
		},
		{
			"seven",
			IncBy[[]float64](7, 1.1),
			30.8,
			true,
		},
		{
			"eight",
			IncBy[[]float64](8, 1.1),
			39.6,
			true,
		},
		{
			"nine",
			IncBy[[]float64](9, 1.1),
			49.5,
			true,
		},
		{
			"ten",
			IncBy[[]float64](10, 1.1),
			60.5,
			true,
		},
		{
			"thirty one",
			Inc[[]float64](31),
			496,
			false,
		},
		{
			"thirty two",
			Inc[[]float64](32),
			528,
			false,
		},
		{
			"thirty three",
			Inc[[]float64](33),
			561,
			false,
		},
		{
			"ninety nine",
			Inc[[]float64](99),
			4950,
			false,
		},
		{
			"hundred",
			Inc[[]float64](100),
			5050,
			false,
		},
		{
			"hundred and one",
			Inc[[]float64](101),
			5151,
			false,
		},
		{
			"thousand",
			Inc[[]float64](1000),
			500500,
			false,
		},
		{
			"ten thousand",
			Inc[[]float64](10000),
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
					return mathx.IsClose(tt.wantRet, gotRet)
				}
				return gotRet == tt.wantRet
			}(); !equal {
				t.Errorf("Sum() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
