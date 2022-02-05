package main

func removeDuplicates(nums []int) int {
	found := make(map[int]bool)

	idx := 0

	var num int
	var ok bool
	for _, num = range nums {
		if _, ok = found[num]; !ok {
			found[num] = true
			nums[idx] = num
			idx++
		}
	}

	return idx
}
