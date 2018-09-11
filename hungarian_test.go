package hungarian_test

import (
	"testing"
	"hungarian"
)

var tests = []struct {
	m      [][]float64
	result map[int]map[int]float64
}{
	{[][]float64{
		{6, 2, 3, 4, 5},
		{3, 8, 2, 8, 1},
		{9, 9, 5, 4, 2},
		{6, 7, 3, 4, 3},
		{1, 2, 6, 4, 9},
	}, map[int]map[int]float64{
		0: {2: 3},
		1: {3: 8},
		2: {0: 9},
		3: {1: 7},
		4: {4: 9},
	}},
}

func TestSolveMax(t *testing.T) {
	for _, value := range tests {
		for key, val := range hungarian.SolveMax(value.m) {
			for k, v := range val {
				if v != value.result[key][k] {
					t.Fatalf("Want %d, got: %d", v, value.result[key][k])
				}
			}
		}
	}
}
