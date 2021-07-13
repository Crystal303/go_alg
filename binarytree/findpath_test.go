package binarytree

import "testing"

func TestFindPath(t *testing.T) {
	tree := initTree1()
	FindPath(tree, 25)
	FindPath(tree, 15)
	FindPath(tree, 8)

}

func ExampleFindPath() {
	tree := initTree1()
	FindPath(tree, 25)

	// output:
	// [8 8 9]
	// [8 8 2 7]
}
