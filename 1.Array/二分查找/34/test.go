package main

import "fmt"

func main() {
	slice := []int{1, 3, 3, 3, 5, 6}
	result := searchRange(slice, 2)
	fmt.Println(result)
}
