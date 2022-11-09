package main

func mySqrt(x int) int {
	low, high := 0, x
	var result int
	for low <= high {
		mid := (low + high) >> 1
		if mid*mid <= x {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
