package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// func insertIntoBST(root *TreeNode, val int) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{
// 			Val:   val,
// 			Left:  nil,
// 			Right: nil,
// 		}
// 	}
// 	var pre *TreeNode
// 	var travel func(node *TreeNode, val int)
// 	travel = func(node *TreeNode, val int) {
// 		if node == nil {
// 			if val > pre.Val {
// 				pre.Right = &TreeNode{
// 					Val:   val,
// 					Left:  nil,
// 					Right: nil,
// 				}
// 			} else {
// 				pre.Left = &TreeNode{
// 					Val:   val,
// 					Left:  nil,
// 					Right: nil,
// 				}
// 			}
// 			return
// 		}
// 		pre = node
// 		if node.Val > val {
// 			travel(node.Left, val)
// 		}
// 		if node.Val < val {
// 			travel(node.Right, val)
// 		}
// 	}
// 	travel(root, val)
// 	return root
// }
