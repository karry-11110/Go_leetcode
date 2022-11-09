package main

import "fmt"
func main() {
	fmt.Print(f16("....~~~~~!!!!"))
}
func f16(str string) int {
	m := map[byte]int{}
	for i := 0; i < len(str); i++{
		if str[i] == '.' || str[i] == '~' ||str[i] == '!' ||str[i] == '@' ||str[i] == '#' ||str[i] == '_' {
			m[str[i]]++
		} 
	}
	n := 0
	for _, v:= range m {
		if v > 1 {
			n++
		}
	}
	return -n*3
}
