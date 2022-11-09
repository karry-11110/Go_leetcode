package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 5, 5, 6}
	result := searchInsert(slice, 5)
	fmt.Println(result)
}
