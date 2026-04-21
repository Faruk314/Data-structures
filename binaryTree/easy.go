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

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	s := stack.Stack[*TreeNode]{}
	s.Push(root)

	for !s.IsEmpty() {
		node, _ := s.Pop()

		result = append(result, node.Val)

		if node.Right != nil {
			s.Push(node.Right)
		}

		if node.Left != nil {
			s.Push(node.Left)
		}
	}

	return result
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	s := stack.Stack[*TreeNode]{}
	var lastVisited *TreeNode
	curr := root

	for !s.IsEmpty() || curr != nil {
		if curr != nil {
			s.Push(curr)
			curr = curr.Left
		} else {

			peekNode, _ := s.Peek()

			if peekNode.Right != nil && lastVisited != peekNode.Right {
				curr = peekNode.Right
			} else {
				result = append(result, peekNode.Val)
				lastVisited, _ = s.Pop()
			}
		}
	}
	return result
}

func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	s := stack.Stack[*TreeNode]{}
	s.Push(root)

	for !s.IsEmpty() {
		node, _ := s.Pop()

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			s.Push(node.Left)
		}

		if node.Right != nil {
			s.Push(node.Right)
		}

	}

	return root
}
