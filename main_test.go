package main

import (
	"strconv"
	"testing"
)

func TestCountTriplets(t *testing.T) {
	testData := []struct {
		arr           []int
		expectedCount int
	}{
		{[]int{1, 5, 3, 2}, 2},
		{[]int{3, 2, 7}, 0},
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{1, 3, 5, 7, 9, 4}, 4},
		{[]int{1, 3, 4, 15, 19}, 2},
		{[]int{7, 2, 5, 4, 3, 6, 1, 9, 10, 12}, 18},
	}

	for _, data := range testData {
		actualCount := CountTriplets(data.arr)
		if actualCount == data.expectedCount {
			t.Log("SUCCESS")
		} else {
			t.Error("Error, expected " + strconv.Itoa(data.expectedCount) + ", got " + strconv.Itoa(actualCount))
		}
	}
}
