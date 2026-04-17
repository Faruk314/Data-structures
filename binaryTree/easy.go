package binarytree

import "golang/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	var result []int
	s := stack.Stack[*TreeNode]{}
	curr := root

	for curr != nil || !s.IsEmpty() {
		for curr != nil {
			s.Push(curr)
			curr = curr.Left
		}

		node, _ := s.Pop()

		result = append(result, node.Val)

		curr = node.Right

	}

	return result
}
