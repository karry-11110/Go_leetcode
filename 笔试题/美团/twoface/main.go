package main

import (
	"fmt" 
	"strings"
	"strconv"
	"sort"
)
func main() {
	str := []string{"10.0.200","11.11","11.10.21"}
	sort.Slice(str, func(i, j int) bool {
		v1 := strings.Split(str[i], ".")
		v2 := strings.Split(str[j], ".")
		for i := 0; i < len(v1) || i < len(v2); i++ {
			x, y := 0, 0
			if i < len(v1) {
				x, _ = strconv.Atoi(v1[i])
			}
			if i < len(v2) {
				y, _ = strconv.Atoi(v2[i])
			}
			if x < y {
				return true
			}
			if x > y {
				return false
			}
		}
		return false
	})
	fmt.Println(str)
}

