package main

func removeElement(nums []int, val int) int {
	numLen := len(nums)

	if numLen < 1 {
		return 0
	}

	i, j := 0, 0

	for ; j < numLen; j++ {
		if nums[j] == val {
			continue
		}
		nums[i] = nums[j]
		i++
	}

	return i
}
