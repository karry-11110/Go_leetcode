package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	//终止条件
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	//找到当前步骤的根节点
	nowRootValue := postorder[len(postorder)-1]
	//从中序遍历中找到一分为二的点的位置，左边为左子树，右边为右子树d
	index := findNowRootIndex(inorder, nowRootValue)
	//构造当前根节点
	nowRoot := &TreeNode{
		Val:   nowRootValue,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
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
