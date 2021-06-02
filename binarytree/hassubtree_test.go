package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasSubTree(t *testing.T) {
	ret := HasSubTree(initTree1(), initTree2())
	assert.True(t, ret)
}

func initTree1() *binaryTreeNode {
	return &binaryTreeNode{
		val: 8,
		left: &binaryTreeNode{
			val: 8,
			left: &binaryTreeNode{
				val:   9,
				left:  nil,
				right: nil,
			},
			right: &binaryTreeNode{
				val: 2,
				left: &binaryTreeNode{
					val:   4,
					left:  nil,
					right: nil,
				},
				right: &binaryTreeNode{
					val:   7,
					left:  nil,
					right: nil,
				},
			},
		},
		right: &binaryTreeNode{
			val:   7,
			left:  nil,
			right: nil,
		},
	}
}

func initTree2() *binaryTreeNode {
	return &binaryTreeNode{
		val: 8,
		left: &binaryTreeNode{
			val:   9,
			left:  nil,
			right: nil,
		},
		right: &binaryTreeNode{
			val:   2,
			left:  nil,
			right: nil,
		},
	}
}
