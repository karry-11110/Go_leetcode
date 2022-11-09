package main

import "fmt"

func main() {
	slice := []int{1, 2, 6, 8}
	result := search(slice, 9)
	fmt.Println(result)
}
