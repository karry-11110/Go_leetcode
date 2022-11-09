package main

func isValid(s string) bool {
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	if s == " " {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
