package vec

import (
	"testing"

	"github.com/razor-87/mathx"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		v, w []float64
		want []float64
	}{
		{
			"2D",
			[]float64{1.1, 2.1},
			[]float64{3.1, 4.1},
			[]float64{4.2, 6.2},
		},
		{
			"2D and 3D",
			[]float64{1.1, 2.1},
			[]float64{4.1, 5.1, 6.1},
			[]float64{5.2, 7.2},
		},
		{
			"3D",
			[]float64{1.1, 2.1, 3.1},
			[]float64{4.1, 5.1, 6.1},
			[]float64{5.2, 7.2, 9.2},
		},
		{
			"3D and 2D",
			[]float64{1.1, 2.1, 3.1},
			[]float64{4.1, 5.1},
			[]float64{5.2, 7.2, 3.1},
		},
		{
			"4D",
			[]float64{1.1, 2.1, 3.1, 4.1},
			[]float64{5.1, 6.1, 7.1, 8.1},
			[]float64{6.2, 8.2, 10.2, 12.2},
		},
		{
			"10D",
			[]float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1},
			[]float64{11.1, 12.1, 13.1, 14.1, 15.1, 16.1, 17.1, 18.1, 19.1, 20.1},
			[]float64{12.2, 14.2, 16.2, 18.2, 20.2, 22.2, 24.2, 26.2, 28.2, 30.2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.v, tt.w)
			for i := range tt.v {
				if !mathx.IsClose(tt.v[i], tt.want[i]) {
					t.Errorf("Add() = %v, want %v", tt.v, tt.want)
				}
			}
		})
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		name string
		v    []float64
		c    float64
		want []float64
	}{
		{
			name: "zeros",
			c:    1,
		},
		{
			"ones",
			[]float64{1, 1, 1},
			1,
			[]float64{1, 1, 1},
		},
		{
			"negation",
			[]float64{1.1, 1.2, 1.3},
			-1,
			[]float64{-1.1, -1.2, -1.3},
		},
		{
			"scaled",
			[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9},
			1.1,
			[]float64{1.21, 2.42, 3.63, 4.84, 6.05, 7.26, 8.47, 9.68, 10.89},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Scale(tt.v, tt.c)
			for i := range tt.v {
				if !mathx.IsClose(tt.v[i], tt.want[i]) {
					t.Errorf("Scale() = %v, want %v", tt.v, tt.want)
				}
			}
		})
	}
}

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
