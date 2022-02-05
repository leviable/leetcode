package main

import (
	"context"
	"sync"
)

const IntNil = 1_000_000_000

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := make([]int, len(nums1)+len(nums2))

	s1 := numStream(ctx, nums1)
	s2 := numStream(ctx, nums2)

	var ok bool
	n1 := IntNil
	n2 := IntNil

	for x := 0; x < len(nums); x++ {
		if n1 == IntNil {
			if n1, ok = <-s1; !ok {
				n1 = IntNil
			}
		}

		if n2 == IntNil {
			if n2, ok = <-s2; !ok {
				n2 = IntNil
			}
		}

		if n1 < n2 {
			nums[x] = n1
			n1 = IntNil
		} else {
			nums[x] = n2
			n2 = IntNil
		}
	}

	median := float64(nums[len(nums)/2])

	if len(nums)%2 == 0 {
		median = (median + float64(nums[(len(nums)/2)-1])) / 2
	}

	return median
}

func numStream(ctx context.Context, nums []int) <-chan int {
	c := make(chan int, 0)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(nums []int, c chan<- int) {
		defer close(c)
		wg.Done()
		for _, n := range nums {
			select {
			case c <- n:
			case <-ctx.Done():
				return
			}
		}
	}(nums, c)

	wg.Wait()

	return c
}
