package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type listnode struct {
	Val  int
	Next *listnode
}
func numstolist(nums []int) *listnode {
	if len(nums) == 0 {
		return nil
	}
	h := &listnode{}
	p := h
	for _, v := range nums{
		p.Next = &listnode{
			Val:  v,
		}
		p = p.Next
	}
	return h.Next

}
func main() {
	//*********创建输入数组

	//[1,2,3,4,5]形式
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//strs := strings.Split(input.Text()," ")
	//nums := make([]int, len(strs))
	//for index, v := range strs{
	//	nums[index], _ = strconv.Atoi(v)
	//}
	//h := numstolist(nums)

	//多组[1,2,3,4,5]形式
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		strs := strings.Split(input.Text()," ")
		nums := make([]int, len(strs))
		for index, v := range strs{
			nums[index], _ = strconv.Atoi(v)
		}
		h := numstolist(nums)
		//********算法处理逻辑
		//********结果输出逻辑
		for h!= nil {
			fmt.Printf("%d ", h.Val)
			h = h.Next
		}
	}


	//********算法处理逻辑
	//********结果输出逻辑
	//for h!= nil {
	//	fmt.Printf("%d ", h.Val)
	//	h = h.Next
	//}
}
