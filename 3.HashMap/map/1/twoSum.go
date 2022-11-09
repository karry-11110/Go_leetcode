package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	result := []int{}
	for index, value := range nums {
		if preIndex, ok := m[target-value]; ok {
			result = []int{index, preIndex}
			break
		} else {
			m[value] = index
		}
	}

	return result
}
