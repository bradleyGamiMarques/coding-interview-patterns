package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
		target  int
	}{
		{name: "test a two element array", numbers: []int{-1, 1}, want: []int{1, 2}, target: 0},
		{name: "test example 1", numbers: []int{2, 7, 11, 15}, want: []int{1, 2}, target: 9},
		{name: "test example 2", numbers: []int{2, 3, 4}, want: []int{1, 3}, target: 6},
		{name: "test example 3", numbers: []int{-1, 0}, want: []int{1, 2}, target: -1},
		{name: "test if both numbers that sum to target are negative", numbers: []int{-3, -2, -1}, want: []int{1, 2}, target: -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
