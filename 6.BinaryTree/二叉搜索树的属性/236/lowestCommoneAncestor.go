package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 前序递归 遍历整棵树找到一种情况（二叉搜索树的最近公共祖先就是中间值）
	if root == nil {
		return root
	}
	//中 ：这里中就是判断当前节点的值与给的两个节点值的关系
	//左
	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil {
			return left
		}
	}
	//右
	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	return root
}
