var pre *TreeNode
var head *TreeNode

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}
	track(pRootOfTree)
	head.Left, pre.Right = pre, head
	return head
}
func track(node *TreeNode) {
	if node == nil {
		return
	}
	track(node.Left)
	if pre == nil {
		head = node
	} else {
		pre.Right, node.Left = node, pre
	}
	pre = node
	track(node.Right)
}