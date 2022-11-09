package main

func isPerfectSquare(num int) bool {
	low, high := 1, num
	var result bool
	for low <= high {
		mid := (low + high) >> 1
		if mid*mid == num {
			result = true
			break
		} else if mid*mid > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}
