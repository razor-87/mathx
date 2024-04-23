package vec

import (
	"testing"

	"github.com/razor-87/mathx"
)

func TestZeros(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Zeros[[]float64](5)
		if len(v) != 5 {
			t.Errorf("expected length 5, got %d", len(v))
		}
		for _, x := range v {
			if x != 0 {
				t.Errorf("expected all elements to be 0, got %f", x)
			}
		}
	})
}

func TestOnes(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Ones[[]float64](5)
		if len(v) != 5 {
			t.Errorf("expected length 5, got %d", len(v))
		}
		for _, x := range v {
			if x != 1 {
				t.Errorf("expected all elements to be 1, got %f", x)
			}
		}
	})
}

func TestRep(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Rep[[]float64](5, 3.14)
		if len(v) != 5 {
			t.Errorf("expected length 5, got %d", len(v))
		}
		for _, x := range v {
			if x != 3.14 {
				t.Errorf("expected all elements to be 3.14, got %f", x)
			}
		}
	})
}

func TestInc(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Inc[[]float64](5)
		if len(v) != 5 {
			t.Errorf("expected length 5, got %d", len(v))
		}
		for i, x := range v {
			if x != float64(i+1) {
				t.Errorf("expected %d, got %f", i+1, x)
			}
		}
	})
}

func TestIncBy(t *testing.T) {
	nums := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 11}
	t.Run("fast", func(t *testing.T) {
		v := IncBy[[]float64](10, 1.1)
		if len(v) != 10 {
			t.Errorf("expected length 10, got %d", len(v))
		}
		for i := range v {
			if !mathx.IsClose(nums[i], v[i]) {
				t.Errorf("expected %f, got %f", nums[i], v[i])
			}
		}
	})
}

func TestRand(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		v := Rand[[]float64](5)
		if len(v) != 5 {
			t.Errorf("expected length 5, got %d", len(v))
		}
		for _, x := range v {
			if x < 0 || x > 1 {
				t.Errorf("expected value between 0 and 1, got %f", x)
			}
		}
	})
}

func TestCopy(t *testing.T) {
	t.Run("fast", func(t *testing.T) {
		src := Inc[[]float64](5)
		v := Copy(src)
		if len(v) != len(src) {
			t.Fatalf("expected length %d, got %d", len(src), len(v))
		}
		if i := 4; v[i] != src[i] {
			t.Errorf("expected %f, got %f", src[i], v[i])
		}
	})
	if !t.Failed() {
		t.Run("deep", func(t *testing.T) {
			src := Ones[[]float64](5)
			v := Copy(src)
			FillZeroes(src)
			for i := range src {
				if src[i] == v[i] {
					t.Error("non-zero value is expected")
				}
			}
		})
	}
}
