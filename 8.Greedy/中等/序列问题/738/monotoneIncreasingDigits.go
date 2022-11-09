package main

import "strconv"

func check(number int) bool {
	i := number % 10
	number /= 10
	for number > 0 {
		j := number % 10
		if j < i {
			return false
		}
		i = j
		number /= 10
	}
	return true
}
func monotoneIncreasingDigits(n int) int {
	//转换为byte型数组
	s := []byte(strconv.Itoa(n))
	length := len(s)
	if length <= 1 {
		return n
	}
	for i := length - 1; i > 0; i-- {
		if s[i-1] > s[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			s[i-1] -= 1
			for j := i; j < length; j++ { //后面的全部置为9
				s[j] = '9'
			}
		}

	}
	result, _ := strconv.Atoi(string(s))
	return result
}
