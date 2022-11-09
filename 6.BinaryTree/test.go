package main

//单链表结点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//构造二叉树，返回根节点
var NULL = -1 << 63 // NULL 方便添加测试数据
func NewBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := make([]*TreeNode, 1, 2*n)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func main() {
	sortedArrayToBST([]int{1, 2, 3})
}
