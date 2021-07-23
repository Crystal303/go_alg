package binarytree

func convert(root *binaryTreeNode) *binaryTreeNode {
	if root == nil {
		return nil
	}

	// 未分配地址 需要拿到返回值
	var lastNode *binaryTreeNode
	lastNode = convertNode(root, lastNode)

	if lastNode == nil {
		return lastNode
	}

	headNode := lastNode
	for headNode.left != nil {
		headNode = headNode.left
	}

	return headNode
}

func convertNode(pNode *binaryTreeNode, lastNode *binaryTreeNode) *binaryTreeNode {
	if pNode == nil {
		return lastNode
	}

	curr := pNode
	if curr.left != nil {
		lastNode = convertNode(curr.left, lastNode)
	}
	curr.left = lastNode
	if lastNode != nil {
		lastNode.right = curr
	}
	lastNode = curr
	if curr.right != nil {
		lastNode = convertNode(curr.right, lastNode)
	}
	return lastNode
}
