package main

func reverseWords(s string) string {
	byteS := []byte(s)
	//双指针思想去除多余空格
	fast, slow := 0, 0
	for byteS[fast] == ' ' && fast < len(byteS) && len(byteS) > 0 {
		fast++
	}
	for ; fast < len(byteS); fast++ {
		if byteS[fast] == ' ' && byteS[fast-1] == byteS[fast] && fast > 1 {
			continue
		}
		byteS[slow] = byteS[fast]
		slow++
	}
	if slow > 1 && byteS[slow-1] == ' ' {
		byteS = byteS[:slow-1]
	} else {
		byteS = byteS[:slow]
	}
	//反转链表
	reverse(byteS, 0, len(byteS)-1)
	for i := 0; i < len(byteS); {
		j := i
		for ; j < len(byteS) && byteS[j] != ' '; j++ {

		}
		reverse(byteS, i, j-1)
		i = j
		i++
	}
	return string(byteS)

}
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
}
