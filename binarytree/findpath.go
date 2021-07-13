package binarytree

import "fmt"

func FindPath(tree *binaryTreeNode, exceptedSum int) {
	if tree == nil {
		return
	}
	path := make([]int, 0)
	findPath(tree, exceptedSum, 0, path)
}

func findPath(tree *binaryTreeNode, exceptedSum, currentSum int, path []int) {
	path = append(path, tree.val)
	currentSum += tree.val

	isLeafNode := false
	if tree.left == nil && tree.right == nil {
		isLeafNode = true
	}

	// 叶子结点且当前和为期待值
	if isLeafNode && currentSum == exceptedSum {
		fmt.Println(path)
	}

	if tree.left != nil {
		findPath(tree.left, exceptedSum, currentSum, path)
	}

	if tree.right != nil {
		findPath(tree.right, exceptedSum, currentSum, path)
	}

	// 回退
	currentSum -= tree.val
	path = path[:len(path)-1]
}
