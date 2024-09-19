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

func TestCosSim(t *testing.T) {
	tests := []struct {
		name    string
		v, w    []float64
		wantRet float64
	}{
		{
			"identical",
			[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10},
			[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.10},
			1,
		},
		{
			"orthogonal",
			[]float64{1.1, 0, 3.3, 0, 5.5, 0, 7.7, 0, 9.9, 0},
			[]float64{0, 2.2, 0, 4.4, 0, 6.6, 0, 8.8, 0, 10.10},
			0,
		},
		{
			"similar",
			[]float64{0.650964, 0.838760, 0.776138, 0.148163, 0.352947, 0.488309, 0.965292, 0.950068, 0.306189, 0.727190, 0.002359, 0.735989, 0.898889, 0.041353, 0.851644, 0.488960, 0.557591, 0.901085, 0.291284, 0.908876, 0.810435, 0.808070, 0.143053},
			[]float64{0.246338, 0.446570, 0.718086, 0.505610, 0.618212, 0.237045, 0.983175, 0.639710, 0.082605, 0.747467, 0.099634, 0.545008, 0.053702, 0.277402, 0.208264, 0.270416, 0.722504, 0.486679, 0.135704, 0.791134, 0.727898, 0.688089, 0.102923},
			0.888489,
		},
		{
			"similar >0.9",
			[]float64{0.329729, 0.598021, 0.367330, 0.258583, 0.466056, 0.608380, 0.453521, 0.687399, 0.513015, 0.766588, 0.344440, 0.035837, 0.792482, 0.874669, 0.301011, 0.601224, 0.623374, 0.483044, 0.023173, 0.758599, 0.710826, 0.155102, 0.266549},
			[]float64{0.521058, 0.459816, 0.211150, 0.594998, 0.634262, 0.988823, 0.609677, 0.754843, 0.749086, 0.318621, 0.789569, 0.417833, 0.809582, 0.788627, 0.403891, 0.365265, 0.826721, 0.643362, 0.311768, 0.967600, 0.367567, 0.586376, 0.090990},
			0.906468,
		},
		{
			"not similar",
			[]float64{0.114759, 0.892615, 0.014377, 0.018034, 0.553304, 0.555431, 0.044975, 0.443808, 0.798458, 0.275288, 0.134466, 0.923047, 0.604047, 0.379725, 0.431604, 0.531235, 0.434814, 0.311153, 0.000953, 0.082129, 0.668995, 0.991486, 0.803755},
			[]float64{0.076804, 0.467502, 0.895549, 0.381623, 0.110903, 0.023189, 0.747739, 0.120228, 0.800729, 0.679519, 0.656822, 0.223147, 0.070646, 0.010816, 0.159206, 0.136339, 0.927702, 0.810313, 0.052294, 0.770802, 0.065279, 0.551639, 0.365337},
			0.563583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := CosSim(tt.v, tt.w); !mathx.IsClose(gotRet, tt.wantRet) {
				t.Errorf("CosSim() = %v, want %v", gotRet, tt.wantRet)
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
