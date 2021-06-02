package binarytree

// 子结构，遍历+递归
func DoesTree1HasTree2(tree1, tree2 *binaryTreeNode) bool {
	if tree2 == nil {
		return true
	}
	if tree1 == nil {
		return false
	}

	if tree1.val == tree2.val {
		return DoesTree1HasTree2(tree1.left, tree2.left) && DoesTree1HasTree2(tree1.right, tree2.right)
	}
	return false
}

func HasSubTree(tree1, tree2 *binaryTreeNode) (ret bool) {
	if tree2 == nil {
		return true
	}
	if tree1 == nil {
		return false
	}

	if tree1.val == tree2.val {
		ret = DoesTree1HasTree2(tree1, tree2)
		if !ret {
			return HasSubTree(tree1.left, tree2) || HasSubTree(tree1.right, tree2)
		}
	}
	return ret
}

