package main

func main() {

}
func partitionLabels(s string) []int {
	result := []int{}
	marks := [26]int{}
	for i := 0; i < len(s); i++ { // 统计每一个字符最后出现的位置
		marks[s[i]-'a'] = i
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, marks[s[i]-'a']) //// 找到字符出现的最远边界
		if i == right {
			result = append(result, (right - left + 1))
			left = i + 1
		}
	}
	return result
}
func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
