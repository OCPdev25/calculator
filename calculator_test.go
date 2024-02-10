package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T){
	t.Parallel()
	type testCase struct {
		fn func (float64, float64) float64
		a, b float64
		want float64
	}
	tcs := []testCase{
		{fn: calculator.Add, a:2, b: 2, want: 4},
		{fn:calculator.Add, a: 1, b: 2, want: 3},
		{fn: calculator.Add, a:0, b: 0, want: 0},
		{fn: calculator.Add, a: -1, b: -1, want: -2},
	}
	for _, tc := range tcs {
		got := calculator.Add(tc.a, tc.b)
	if tc.want != got {
		t.Errorf("Add(%.1f, %.1f): want %f , got %f", tc.a, tc.b, tc.want, got)
	}

	}


}

func TestSubtract(t *testing.T){
	t.Parallel()
	var a, b float64
	a = 2
	b = 2
	var want float64 = 0
	got := calculator.Subtract(2,2)
	if want != got {
		t.Errorf("Subtract(%.1f, %.1f): want %f , got %f", a, b, want, got)
	}
}
