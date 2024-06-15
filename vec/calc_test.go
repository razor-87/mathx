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

func TestAddScaled(t *testing.T) {
	scalar := 1.1
	tests := []struct {
		name string
		v    []float64
		w    []float64
		want []float64
	}{
		{
			"2D",
			[]float64{1.1, 2.1},
			[]float64{3.1, 4.1},
			[]float64{4.51, 6.61},
		},
		{
			"3D",
			[]float64{1.1, 2.1, 3.1},
			[]float64{4.1, 5.1, 6.1},
			[]float64{5.61, 7.71, 9.81},
		},
		{
			"4D",
			[]float64{1.1, 2.1, 3.1, 4.1},
			[]float64{5.1, 6.1, 7.1, 8.1},
			[]float64{6.71, 8.81, 10.91, 13.01},
		},
		{
			"10D",
			[]float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1},
			[]float64{11.1, 12.1, 13.1, 14.1, 15.1, 16.1, 17.1, 18.1, 19.1, 20.1},
			[]float64{13.31, 15.41, 17.51, 19.61, 21.71, 23.81, 25.91, 28.01, 30.11, 32.21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddScaled(tt.v, tt.w, scalar)
			for i := range tt.v {
				if !mathx.IsClose(tt.v[i], tt.want[i]) {
					t.Errorf("AddScaled() = %v, want %v", tt.v, tt.want)
				}
			}
		})
	}
}

func TestDotProd(t *testing.T) {
	tests := []struct {
		name    string
		v, w    []float64
		wantRet float64
	}{
		{
			"equal dimensions",
			[]float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1},
			[]float64{11.1, 12.1, 13.1, 14.1, 15.1, 16.1, 17.1, 18.1, 19.1, 20.1},
			956.1,
		},
		{
			"different dimensions",
			[]float64{1.1, 2.1, 3.1},
			[]float64{4.1, 5.1},
			15.22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := DotProd(tt.v, tt.w); !mathx.IsClose(gotRet, tt.wantRet) {
				t.Errorf("DotProd() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		name    string
		v       []float64
		wantRet float64
	}{
		{
			"2D",
			[]float64{1.1, 2.1},
			2.370653,
		},
		{
			"3D",
			[]float64{1.1, 2.1, 3.1},
			3.902563,
		},
		{
			"4D",
			[]float64{1.1, 2.1, 3.1, 4.1},
			5.660388,
		},
		{
			"10D",
			[]float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1},
			19.902261,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Length(tt.v); !mathx.IsClose(gotRet, tt.wantRet) {
				t.Errorf("Length() = %v, want %v", gotRet, tt.wantRet)
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

func TestUnit(t *testing.T) {
	tests := []struct {
		name string
		v    []float64
		want []float64
	}{
		{
			"2D",
			[]float64{1, 2},
			[]float64{0.447213, 0.894427},
		},
		{
			"3D",
			[]float64{1, 2, 3},
			[]float64{0.267261, 0.534522, 0.801783},
		},
		{
			"4D",
			[]float64{1, 2, 3, 4},
			[]float64{0.182574, 0.365148, 0.547722, 0.730296},
		},
		{
			"10D",
			[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0.050964, 0.101929, 0.152894, 0.203858, 0.254823, 0.305788, 0.356753, 0.407717, 0.458682, 0.509647},
		},
		{
			"one non-zero component",
			[]float64{0, 0, 1, 0, 0},
			[]float64{0, 0, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Unit(tt.v)
			for i := range tt.v {
				if !mathx.IsClose(tt.v[i], tt.want[i]) {
					t.Errorf("Unit() = %v, want %v", tt.v, tt.want)
				}
			}
		})
	}
}
