package mathx

import (
	"math"
	"testing"
)

func TestIsClose(t *testing.T) {
	type args[T Floaty] struct {
		x T
		y T
	}
	tests32 := []struct {
		name string
		args args[float32]
		want bool
	}{
		{
			name: "float32 close",
			args: args[float32]{
				x: 1.0,
				y: 1.000001,
			},
			want: true,
		},
		{
			name: "float32 not close",
			args: args[float32]{
				x: 1.0,
				y: 1.001,
			},
			want: false,
		},
		{
			name: "float32 equal",
			args: args[float32]{
				x: 1.0,
				y: 1.0,
			},
			want: true,
		},
		{
			name: "float32 zero",
			args: args[float32]{
				x: 0.0,
				y: 0.0,
			},
			want: true,
		},
		{
			name: "float32 NaN",
			args: args[float32]{
				x: float32(math.NaN()),
				y: float32(math.NaN()),
			},
			want: true,
		},
		{
			name: "float32 NaN and zero",
			args: args[float32]{
				x: float32(math.NaN()),
				y: 0.0,
			},
			want: false,
		},
	}
	for _, tt := range tests32 {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsClose(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsClose() = %v, want %v", got, tt.want)
			}
		})
	}
	tests64 := []struct {
		name string
		args args[float64]
		want bool
	}{
		{
			name: "float64 close",
			args: args[float64]{
				x: 1.0,
				y: 1.000001,
			},
			want: true,
		},
		{
			name: "float64 not close",
			args: args[float64]{
				x: 1.0,
				y: 1.001,
			},
			want: false,
		},
		{
			name: "float64 equal",
			args: args[float64]{
				x: 1.0,
				y: 1.0,
			},
			want: true,
		},
		{
			name: "float64 zero",
			args: args[float64]{
				x: 0.0,
				y: 0.0,
			},
			want: true,
		},
		{
			name: "float64 NaN",
			args: args[float64]{
				x: math.NaN(),
				y: math.NaN(),
			},
			want: true,
		},
		{
			name: "float64 NaN and zero",
			args: args[float64]{
				x: math.NaN(),
				y: 0.0,
			},
			want: false,
		},
	}
	for _, tt := range tests64 {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsClose(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsClose() = %v, want %v", got, tt.want)
			}
		})
	}
}
