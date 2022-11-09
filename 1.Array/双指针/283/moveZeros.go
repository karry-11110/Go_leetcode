package main

func moveZeroes(nums []int) {
	length := len(nums)
	slow := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}
	for i := slow; i < length; i++ {
		nums[i] = 0
	}

}
