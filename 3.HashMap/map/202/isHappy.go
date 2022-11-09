package main

func isHappy(n int) bool {
	a := map[int]bool{}
	for n != 1 && !a[n] {
		n, a[n] = getSum(n), true
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
