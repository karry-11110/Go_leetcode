func minNumber(nums []int) string {
	var result string
	if len(nums) == 0 {
		return result
	}
	strs := make([]string, 0)
	for i := range nums {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	QuickSort(strs, 0, len(strs)-1)
	for i := range strs {
		result += strs[i]
	}
	return result
}
func QuickSort(strs []string, begin, end int) {
	if begin < end {
		loc := partition(strs, begin, end)
		QuickSort(strs, begin, loc-1)
		QuickSort(strs, loc+1, end)
	}
}
func partition(strs []string, begin, end int) int {
	i, j := begin+1, end
	for i < j {
		if (strs[begin] + strs[i]) < (strs[i] + strs[begin]) {
			strs[i], strs[j] = strs[j], strs[i]
			j--
		} else {
			i++
		}
	}
	if (strs[begin] + strs[i]) <= (strs[i] + strs[begin]) {
		i--
	}
	strs[begin], strs[i] = strs[i], strs[begin]
	return i
}