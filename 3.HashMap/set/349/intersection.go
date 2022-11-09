package main

//set实现
func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]struct{}{} //go的set实现，空结构体不占内存
	result := []int{}
	for _, value := range nums1 {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
		}
	}
	for _, value := range nums2 {
		if _, ok := set[value]; ok {
			result = append(result, value)
			delete(set, value) //用完就甩掉了
		}
	}
	return result
}

// func intersection(nums1 []int, nums2 []int) []int {
// 	m := map[int]int{}
// 	result := []int{}
// 	for _, value := range nums1 {
// 		m[value] = 1
// 	}
// 	for _, value := range nums2 {
// 		if count, ok := m[value]; ok && count > 0 {
// 			result = append(result, value)
// 			m[value]--
// 		}
// 	}
// 	return result
// }
