//直接选择排序
func SelectSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代，最后一轮自动号
	for i := 0; i < n-1; i++ {
		// 每次从第 i 位开始，找到最小的元素
		min := list[i] // 最小数
		minIndex := i  // 最小数的下标
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				// 如果找到的数比上次的还小，那么最小的数变为它
				min = list[j]
				minIndex = j
			}
		}
		// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}