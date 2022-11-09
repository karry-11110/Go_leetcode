package main

import "sort"

//有序情况下使用双指针
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	intersection := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersection = append(intersection, nums1[index1])
			index1++
			index2++
		}
	}
	return intersection
}

// //哈希map表
// func intersect(nums1 []int, nums2 []int) []int {
// 	result := []int{}
// 	a1 := map[int]int{}
// 	for _, value := range nums1 {
// 		a1[value]++
// 	}
// 	for _, value := range nums2 {
// 		if a1[value] > 0 {
// 			result = append(result, value)
// 			a1[value]--
// 		}
// 	}
// 	return result
// }
