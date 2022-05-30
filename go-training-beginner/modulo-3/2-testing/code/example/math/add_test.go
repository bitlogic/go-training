package math_test

import (
	"example/math"
	"testing"
)

func TestAdd(t *testing.T) {
	got := math.Add(1, 1)
	expected := 2

	if got != expected {
		t.Fail()
	}
}

//------//
//Table driven tests
//------//

type addTestCase struct {
	a, b, expected int
}

var testCases = []addTestCase{
	{1, 1, 2},
	{25, 25, 50},
	{2, 1, 3},
	{1, 10, 11},
}

func TestAddTDT(t *testing.T) {

	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)

		if got != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, got)
		}
	}
}

func FuzzTestAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		math.Add(a, b)
	})
}

func FuzzTestAddMore(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		math.AddMore(a, b)
	})
}
