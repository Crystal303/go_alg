package binarytree

// 给定一个二叉树，返回该二叉树的镜像
// 非递归
func MirrorNonRecursively(header *binaryTreeNode) {
	if header == nil {
		return
	}
	list := make([]*binaryTreeNode, 0)
	list = append(list, header)

	for len(list) > 0 {
		curr := list[0]
		list = list[1:]

		if curr.left != nil ||
			curr.right != nil {
			temp := curr.left
			curr.left = curr.right
			curr.right = temp

			if curr.left != nil {
				list = append(list, curr.left)
			}
			if curr.right != nil {
				list = append(list, curr.right)
			}
		}
	}
}

func MirrorRecursively(header *binaryTreeNode) {
	if (header == nil) ||
		(header.left == nil && header.right == nil) {
		return
	}

	temp := header.left
	header.left = header.right
	header.right = temp

	if header.left != nil {
		MirrorRecursively(header.left)
	}
	if header.right != nil {
		MirrorRecursively(header.right)
	}
}
