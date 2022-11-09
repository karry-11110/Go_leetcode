func countDigitOne(n int) int {
	base, result := 1, 0
	for n/base != 0 {
		curr := n / base % 10
		a, b := n/base/10, n%base
		if curr == 1 {
			result += a*base + b + 1
		} else if curr > 1 {
			result += (a + 1) * base
		} else {
			result += a * base
		}
		base *= 10
	}
	return result
}
