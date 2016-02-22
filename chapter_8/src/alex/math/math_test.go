package math

import "testing"

type testpair struct {
	values          []float64
	expected_result float64
}

var tests_avg = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests_avg {
		average := Average(pair.values)
		if average != pair.expected_result {
			t.Error("Expected", pair.expected_result, ", got", average)
		}
	}
}

var tests_max = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{2, 1}, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
	{[]float64{-3, -1}, -1},
}

func TestMax(t *testing.T) {
	for _, pair := range tests_max {
		max := Max(pair.values)
		if max != pair.expected_result {
			t.Error("Expected", pair.expected_result, ", got", max)
		}
	}
}

var tests_min = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{2, 1}, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, -1},
	{[]float64{-3, -1}, -3},
}

func TestMin(t *testing.T) {
	for _, pair := range tests_min {
		min := Min(pair.values)
		if min != pair.expected_result {
			t.Error("Expected", pair.expected_result, ", got", min)
		}
	}
}
