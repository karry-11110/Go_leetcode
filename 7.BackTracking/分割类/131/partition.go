package main

func partition(s string) [][]string {
	var result [][]string
	var path []string
	var backtracking func(s string, startIndex int)
	backtracking = func(s string, startIndex int) {
		//切割线的位置就是startIndex的位置
		if startIndex >= len(s) {
			result = append(result, append([]string{}, path...))
			return
		}
		for i := startIndex; i < len(s); i++ {
			//定义了起始位置startIndex，那么 [startIndex, i] 就是要截取的子串,首先判断这个子串是不是回文，
			//如果是回文，就加入在vector<string> path中，path用来记录切割过的回文子串。
			if isPalindrome(s, startIndex, i) {
				path = append(path, s[startIndex:i+1])
			} else {
				continue
			}
			//递归
			backtracking(s, i+1) //寻找i+1为起始位置的子串
			//回溯
			path = path[:len(path)-1]

		}
	}
	backtracking(s, 0)
	return result
}

//判断是否为回文
func isPalindrome(s string, startIndex, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}
