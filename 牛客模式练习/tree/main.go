package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type treenode struct {
	Val  int
	Left *treenode
	Right *treenode
}
func numstotree(nums []int, index int) *treenode {
	if index >= len(nums) || nums[index] == -1 {
		return nil
	}
	root := &treenode{
		Val: nums[index],
		Left: numstotree(nums, index*2 + 1),
		Right: numstotree(nums, index*2 + 2),
	}
	return root

}
func  printtree(node *treenode)  {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	printtree(node.Left)
	printtree(node.Right)
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
	//h := numstotree(nums, 0)

	//多组[1,2,3,4,5]形式
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		strs := strings.Split(input.Text()," ")
		nums := make([]int, len(strs))
		for index, v := range strs{
			nums[index], _ = strconv.Atoi(v)
		}
		h := numstotree(nums, 0)
		//********算法处理逻辑
		//********结果输出逻辑
		printtree(h)
	}

	//********算法处理逻辑
	//********结果输出逻辑
	//printtree(h)
}
