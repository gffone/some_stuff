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

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
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

func isSame(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			if queue[i] == nil && queue[qLen-1] == nil {
				continue
			}
			if queue[i] != nil || queue[qLen-1] != nil {
				return false
			}
			if queue[i].Val != queue[qLen-1].Val {
				return false
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		queue = queue[qLen:]
	}
	return true
}

// build tree from slice
func buildTree(slice []int, i int) *TreeNode {
	if i >= len(slice) {
		return nil
	}

	node := newTreeNode(slice[i])

	node.Left = buildTree(slice, 2*i+1)
	node.Right = buildTree(slice, 2*i+2)

	return node
}
