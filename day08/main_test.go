package main

import "testing"

func Test(t *testing.T) {
	slice5 := []string{"a", "b", "d", "f", "g"}
	slice6 := []string{"a", "b", "d", "e", "f", "g"}

	if len(sliceDifference(slice5, slice6)) != 1 {
		t.Error("Expected 1, got ", len(sliceDifference(slice5, slice6)))
	}
}
