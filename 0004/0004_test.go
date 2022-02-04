package main

import (
	"fmt"
	"testing"
)

func TestFindMedian(t *testing.T) {
	examples := []struct {
		nums1, nums2 []int
		want         float64
	}{
		{nums1: []int{1, 3}, nums2: []int{2}, want: 2.0},
		{nums1: []int{2}, nums2: []int{1, 3}, want: 2.0},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.5},
		{nums1: []int{3, 4}, nums2: []int{1, 2}, want: 2.5},
		{nums1: []int{}, nums2: []int{1}, want: 1.0},
		{nums1: []int{1}, nums2: []int{}, want: 1.0},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("%v %v", tt.nums1, tt.nums2), func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			want := tt.want

			if got != want {
				t.Errorf("got %f, want %f", got, want)
			}
		})
	}
}
