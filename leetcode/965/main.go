package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && (root.Left.Val != root.Val) {
		return false
	}
	if root.Right != nil && (root.Right.Val != root.Val) {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
