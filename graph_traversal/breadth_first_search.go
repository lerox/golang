package graph_traversal

import (
	"container/list"
	"fmt"
)

func (bn *BinaryNode) BreadthFirstSearch() {

	queue := list.New()
	queue.PushBack(*bn)

	for queue.Len() != 0 {

		element := queue.Front()
		currentNode := element.Value.(BinaryNode)
		queue.Remove(element)

		fmt.Print(currentNode.Value)

		if currentNode.Left != nil {
			queue.PushBack(*currentNode.Left)
		}

		if currentNode.Right != nil {
			queue.PushBack(*currentNode.Right)
		}
	}
}
