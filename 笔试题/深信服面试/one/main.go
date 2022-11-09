package main

import (
	//"math/rand"
	//"time"
	"fmt"
)
func main() {
	//rand.Seed(time.Now().UnixNano()) // 取纳秒时间戳，可以保证每次的随机数种子都不同
	//for i := 0; i < 5; i++ {
	//	row, col := rand.Intn(100),rand.Intn(100)
	//	fmt.Println(row, col)
	//}
	m := 1001//会产生0-1000的随机数
	a, b := 97, 100
	x := 1
	for i := 0; i < 10; i++ {
		x = (a*x+b) % m
		fmt.Println(x)
	}
}
