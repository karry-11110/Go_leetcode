package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	//终止条件
	if len(inorder) < 1 || len(preorder) < 1 {
		return nil
	}
	//找到当前步骤的根节点
	nowRootValue := preorder[0]
	//从中序遍历中找到一分为二的点的位置，左边为左子树，右边为右子树d
	index := findNowRootIndex(inorder, nowRootValue)
	//构造当前根节点
	nowRoot := &TreeNode{
		Val:   nowRootValue,
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return nowRoot
}
func findNowRootIndex(inorder []int, nowRootValue int) int {
	for i := 0; i < len(inorder); i++ {
		if nowRootValue == inorder[i] {
			return i
		}
	}
	return -1
}
