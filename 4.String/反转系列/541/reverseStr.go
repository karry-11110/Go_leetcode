package main

func reverseStr(s string, k int) string {
	byteS := []byte(s)
	for i := 0; i < len(byteS); i += 2 * k {
		j := i + k - 1
		if j <= len(byteS)-1 {
			reverse(byteS, i, j)
		} else {
			reverse(byteS, i, len(byteS)-1)
		}
	}
	return string(byteS)
}
func reverse(s []byte, i, j int) {
	left, right := i, j
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
