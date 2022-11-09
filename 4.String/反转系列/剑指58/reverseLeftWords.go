package main

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 1.  反转整个字符
	// 2. 反转前len，len-n个字符
	// 3.反转第len-len-n到end字符
	reverse(b, 0, len(b)-1)
	reverse(b, 0, len(b)-n-1)
	reverse(b, len(b)-n, len(b)-1)

	return string(b)
}

// 切片是引用传递
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// //将前n个元素添加到tmp集合中，然后在ss[n:length]集合后边依次添加tmp集合中的元素即可
// func reverseLeftWords(s string, n int) string {
// 	ss := []byte(s)
// 	length := len(ss)
// 	var tmp []byte
// 	for i := 0; i < n; i++ {
// 		tmp = append(tmp, ss[i])
// 	}
// 	ss = append(ss[n:length], tmp...)
// 	return string(ss)
// }
