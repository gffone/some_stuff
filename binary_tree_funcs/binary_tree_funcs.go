package binary_tree_funcs

import (
	"fmt"
	"math"
)

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

func isBalanced(t *TreeNode) bool {
	if t == nil {
		return true
	}

	if math.Abs(float64(height(t.Left)-height(t.Right))) > 1 {
		return false
	}

	return isBalanced(t.Left) && isBalanced(t.Right)
}

func DFTLNR(t *TreeNode) {
	if t == nil {
		return
	}

	DFTLNR(t.Left)
	fmt.Print(t.Val)
	DFTLNR(t.Right)
}
