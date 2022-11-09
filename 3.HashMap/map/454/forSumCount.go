package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m := map[int]int{}
	for _, value1 := range nums1 {
		for _, value2 := range nums2 {
			m[value1+value2]++
		}
	}
	for _, value3 := range nums3 {
		for _, value4 := range nums4 {
			count += m[-(value3 + value4)]
		}
	}
	return count
}
