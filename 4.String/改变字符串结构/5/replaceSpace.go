package main

// //原地添加
// func replaceSpace(s string) string {
// 	ss := []byte(s)
// 	result := make([]byte, 0)
// 	for i := 0; i < len(ss); i++ {
// 		if ss[i] == ' ' {
// 			result = append(result, []byte("%20")...)
// 		} else {
// 			result = append(result, ss[i])
// 		}
// 	}
// 	return string(result)
// }

//双指针不添加额外空间
func replaceSpace(s string) string {
	ss := []byte(s)
	//统计空格数量
	count := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			count++
		}
	}
	//拓展原有切片
	length := len(ss)
	tmp := make([]byte, 2*count)
	ss = append(ss, tmp...) //切片添加切片的写法
	//从后往前填充
	i := length - 1
	j := len(ss) - 1
	for i >= 0 {
		if ss[i] != ' ' {
			ss[j] = ss[i]
			i--
			j--
		} else {
			ss[j] = '0'
			ss[j-1] = '2'
			ss[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(ss)

}
