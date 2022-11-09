package main

import "fmt"

func main() {
	slice := []int{1, 3, 0, 5, 7, 9}
	moveZeroes(slice)
	fmt.Println(slice)
}
