package main

func lemonadeChange(bills []int) bool {
	result := true
	left := [2]int{0, 0}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			left[0]++
		}
		if bills[i] == 10 {
			if left[0] >= 1 {
				left[0]--
				left[1]++
			} else {
				return false
			}
		}
		if bills[i] == 20 {
			if left[0] >= 1 && left[1] >= 1 {
				left[0]--
				left[1]--
			} else if left[1] == 0 && left[0] >= 3 {
				left[0] -= 3
			} else {
				return false
			}
		}

	}
	return result
}
