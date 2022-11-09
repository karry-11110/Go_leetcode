package main

func convertBST(root *TreeNode) *TreeNode {
	//反中序遍历,想象成数组该怎么解题，反过来想
	var sum = 0
	travel(root, &sum)
	return root
}
func travel(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	travel(node.Right, sum)
	tmp := *sum
	*sum += node.Val
	node.Val += tmp
	travel(node.Left, sum)

}
