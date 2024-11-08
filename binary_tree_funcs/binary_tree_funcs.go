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

func dftlnr(t *TreeNode) {
	if t == nil {
		return
	}

	dftlnr(t.Left)
	fmt.Print(t.Val)
	dftlnr(t.Right)
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}


test