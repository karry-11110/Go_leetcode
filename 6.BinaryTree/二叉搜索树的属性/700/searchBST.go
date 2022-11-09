package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return nil

}

//迭代
func searchBST_1(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			break
		}

	}
	return root

}
