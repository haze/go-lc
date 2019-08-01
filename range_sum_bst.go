package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	leftSum := rangeSumBST(root.Left, L, R)
	rightSum := rangeSumBST(root.Right, L, R)
	sum := leftSum + rightSum
	if root.Val >= L && root.Val <= R {
		return root.Val + sum
	}
	return sum // don't add our node value
}
