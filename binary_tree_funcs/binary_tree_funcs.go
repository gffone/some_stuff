package binary_tree_funcs

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(t *TreeNode) int {
	if t == nil {
		return 0
	}

	leftHeight := height(t.Left)
	rightHeight := height(t.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if math.Abs(float64(height(root.Left)-height(root.Right))) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
