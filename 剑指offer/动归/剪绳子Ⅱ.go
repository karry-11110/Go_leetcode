//用动态规划会造成长度越界：所以用数学推导会比较好点。
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	sum := 1
	for n > 4 {
		sum *= 3
		sum = sum % 1000000007
		n -= 3
	}
	return sum * n % 1000000007
}