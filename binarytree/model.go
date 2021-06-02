package binarytree

import "fmt"

type binaryTreeNode struct {
	val   int
	left  *binaryTreeNode
	right *binaryTreeNode
}

// 中序
func (b *binaryTreeNode) inPrint() {
	if b.left != nil {
		b.left.inPrint()
	}
	fmt.Print(b.val, " ")
	if b.right != nil {
		b.right.inPrint()
	}
}

func (b *binaryTreeNode) inPrint2() {
	curr := b
	l := make([]*binaryTreeNode, 0)

	for curr != nil || len(l) > 0 {
		for curr != nil {
			l = append(l, curr)
			curr = curr.left
		}
		curr = l[len(l)-1]
		l = l[:len(l)-1]
		fmt.Print(curr.val, " ")
		curr = curr.right
	}
}

// 先序遍历
func (b *binaryTreeNode) prePrint() {
	fmt.Print(b.val, " ")
	if b.left != nil {
		b.left.prePrint()
	}
	if b.right != nil {
		b.right.prePrint()
	}
}

// 先序遍历
func (b *binaryTreeNode) prePrint2() {
	curr := b
	l := make([]*binaryTreeNode, 0)

	for curr != nil || len(l) > 0 {
		for curr != nil {
			fmt.Print(curr.val, " ")
			l = append(l, curr)
			curr = curr.left
		}

		for len(l) > 0 && curr == nil {
			curr = l[len(l)-1].right
			l = l[:len(l)-1]
		}
	}
}

// 后序
func (b *binaryTreeNode) postPrint() {
	if b.left != nil {
		b.left.postPrint()
	}
	if b.right != nil {
		b.right.postPrint()
	}
	fmt.Print(b.val, " ")
}

// 后序(先序改）
func (b *binaryTreeNode) postPrint2() {
	curr := b
	l := make([]*binaryTreeNode, 0)
	pl := make([]int, 0)

	for curr != nil || len(l) > 0 {
		for curr != nil {
			l = append(l, curr)
			pl = append(pl, curr.val)
			curr = curr.right
		}

		for len(l) > 0 && curr == nil {
			curr = l[len(l)-1].left
			l = l[:len(l)-1]
		}
	}
	for i := len(pl) - 1; 0 <= i; i-- {
		fmt.Print(pl[i], " ")
	}
}

// 后序遍历
// 遍历条件
// 1. 当前为叶子节点
// 2. 当前节点的右节点为上一个访问的节点
func (b *binaryTreeNode) postPrint3() {
	var pre *binaryTreeNode
	curr := b
	l := make([]*binaryTreeNode, 0)

	for curr != nil || len(l) > 0 {
		for curr != nil {
			l = append(l, curr)
			curr = curr.left
		}

		for len(l) > 0 && curr == nil {
			curr = l[len(l)-1]
			l = l[:len(l)-1]

			if (curr.left == nil && curr.right == nil) ||
				curr.right == pre {
				fmt.Print(curr.val, " ")
				pre = curr
				curr = nil
			} else {
				l = append(l, curr)
				curr = curr.right
			}
		}
	}
}

// 层次
func (b *binaryTreeNode) levelPrint() {
	l := make([]*binaryTreeNode, 0)
	l = append(l, b)
	for len(l) > 0 {
		node := l[0]
		fmt.Print(node.val, " ")
		if node.left != nil {
			l = append(l, node.left)
		}
		if node.right != nil {
			l = append(l, node.right)
		}
		l = l[1:]
	}
}