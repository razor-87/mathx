package vec

import (
	"testing"
)

func TestZeros(t *testing.T) {
	t.Run("five", func(t *testing.T) {
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
	t.Run("five", func(t *testing.T) {
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
	t.Run("five", func(t *testing.T) {
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
	t.Run("five", func(t *testing.T) {
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
	t.Run("ten", func(t *testing.T) {
		v := IncBy[[]float64](10, 1.1)
		if len(v) != 10 {
			t.Errorf("expected length 10, got %d", len(v))
		}
		for i, x := range v {
			if x != float64(i+1)*1.1 {
				t.Errorf("expected %f, got %f", float64(i+1)*1.1, x)
			}
		}
	})
}

func TestRand(t *testing.T) {
	t.Run("five", func(t *testing.T) {
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
