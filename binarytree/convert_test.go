package binarytree

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	tree := &binaryTreeNode{
		val: 8,
		left: &binaryTreeNode{
			val: 7,
			left: &binaryTreeNode{
				val:  5,
				left: nil,
				right: &binaryTreeNode{
					val:   6,
					left:  nil,
					right: nil,
				},
			},
			right: nil,
		},
		right: &binaryTreeNode{
			val: 10,
			left: &binaryTreeNode{
				val:   9,
				left:  nil,
				right: nil,
			},
			right: &binaryTreeNode{
				val:   11,
				left:  nil,
				right: nil,
			},
		},
	}
	tree.inPrint2()
	fmt.Println()

	doublyLinkedList := convert(tree)
	for doublyLinkedList != nil {
		fmt.Printf("%d ", doublyLinkedList.val)
		doublyLinkedList = doublyLinkedList.right
	}
}

func TestNode(t *testing.T) {
	node := new(binaryTreeNode)
	setNode(node)
	t.Log(node)
}

func setNode(node *binaryTreeNode) {
	node.left = &binaryTreeNode{
		val:   1,
		left:  nil,
		right: nil,
	}
}
