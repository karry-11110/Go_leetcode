package main

import (
	"fmt"
	"math"
)
var res = math.MaxInt64
func main() {
	m := 4
	nums := []int{0,1,2,3}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = solve(nums[i], m, 0)
		fmt.Print(res[i])
		if i != len(nums)-1{
			fmt.Print(" ")
		}
	}
}
func solve(number , m , key int) int {
	curr := number
	if curr % m == 0 {
		return key
	}
	//index := 0
	//distance := math.MaxInt64
	for i := 0; i <= 9 ;i++{
		add := int(math.Pow(10.0, float64(i)))
		if (curr + add ) % m == 0 {
			return key+1
		}
		curr += int(math.Pow(10.0, float64(i)))
		res = min(res, solve(curr, m, key+1))
		//v := (curr + add ) % m
		//if v < distance {
		//	distance = v
		//	index = i
		//}
	}
	return res

}
func min(i, j int) int{
	if i < j {
		return i
	}
	return j
}
