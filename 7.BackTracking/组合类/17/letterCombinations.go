package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 || len(digits) > 4 { //根据题目给出的数据范围
		return nil
	}
	m := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	var result []string
	var s string
	var backtracking func(digits string, index int)
	backtracking = func(digits string, index int) {
		if index == len(digits) {
			result = append(result, s)
			return
		}
		digit := digits[index] - '0' //首先将digits里面对应的第index个字符转换为相应的int型
		letters := m[digit]          //取字符集
		for i := 0; i < len(letters); i++ {
			s = s + string(letters[i]) //处理
			backtracking(digits, index+1)
			s = s[:len(s)-1]
		}
	}
	backtracking(digits, 0)
	return result
}
