package graph_traversal

import "fmt"

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

func NewBinaryNode(value int) *BinaryNode {
	return &BinaryNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (bn *BinaryNode) DepthFirstSearch() {

	var stack Stack

	stack.Push(*bn)

	for !stack.IsEmpty() {

		currentNode := stack.Pop()

		fmt.Print(currentNode.Value)

		if currentNode.Right != nil {
			stack.Push(*currentNode.Right)
		}

		if currentNode.Left != nil {
			stack.Push(*currentNode.Left)
		}
	}

}

type Stack []BinaryNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str BinaryNode) {
	*s = append(*s, str)
}

func (s *Stack) Pop() *BinaryNode {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]

		return &element
	}
}
