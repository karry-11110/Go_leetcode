package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 7, 9}
	result := removeElement(slice, 3)
	fmt.Println(result)
}
