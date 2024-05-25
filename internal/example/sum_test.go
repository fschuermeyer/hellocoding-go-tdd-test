package example_test

import (
	"math"
	"testing"

	"github.com/fschuermeyer/hellocoding-go-tdd-test/internal/example"
)

func TestSum(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{1, 1, 2},
		{120, 240, 360},
		{120, -10, 0},
		{-120, 10, 0},
		{0, 0, 0},
		{math.MaxInt64, -1, 0},
		{9999, 1, 10000},
	}

	for _, c := range cases {
		t.Run("Sum", func(t *testing.T) {
			r := example.Sum(c.a, c.b)

			if r != c.want {
				t.Errorf("TestSum(%d, %d) == %d, want %d", c.a, c.b, r, c.want)
			}
		})
	}
}
