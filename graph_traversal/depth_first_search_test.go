package graph_traversal

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootNode := NewBinaryNode(0)
	rootNode.Left = NewBinaryNode(1)
	rootNode.Right = NewBinaryNode(2)
	rootNode.Left.Left = NewBinaryNode(3)
	rootNode.Left.Right = NewBinaryNode(4)
	rootNode.Right.Left = NewBinaryNode(5)
	rootNode.Right.Right = NewBinaryNode(6)

	// node looks like:
	//     0
	//  1     2
	// 3 4   5 6

	rootNode.DepthFirstSearch()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "0134256"

	if string(out) != expected {
		t.Errorf("We expected %s but we got %s", expected, string(out))
	}
}
