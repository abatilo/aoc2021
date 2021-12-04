package main

import "testing"

// Write tests for gridHasBingo with test list of drawn numbers
func TestGridHasBingo(t *testing.T) {
	grid := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	if !gridHasBingo(grid, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}) {
		t.Errorf("gridHasBingo(%v, %v) = false, want true", grid, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24})
	}
}
