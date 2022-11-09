package main

func sortedSquares(nums []int) []int {
	slow := -1
	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ { //找fast和slow指针如何找？
		if nums[i] < 0 {
			slow = i
		}
	}
	fast := slow + 1
	for fast <= len(nums)-1 && slow >= 0 {
		fastSquares := nums[fast] * nums[fast]
		slowSquares := nums[slow] * nums[slow]
		if fastSquares <= slowSquares {
			result = append(result, fastSquares)
			fast++

		} else {
			result = append(result, slowSquares)
			slow--

		}
	}
	for slow >= 0 {
		result = append(result, nums[slow]*nums[slow])
		slow--

	}
	for fast <= len(nums)-1 {
		result = append(result, nums[fast]*nums[fast])
		fast++

	}
	return result
}

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		if leftValue, rightValue := nums[i]*nums[i], nums[j]*nums[j]; leftValue > rightValue {
			result[pos] = leftValue
			i++
		} else {
			result[pos] = rightValue
			j--
		}
	}
	return result
}
