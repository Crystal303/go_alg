package binarytree

import (
	"fmt"
)

func ExampleMirrorNonRecursively() {
	header := initTree1()
	header.levelPrint()

	fmt.Print("mirror:")
	MirrorNonRecursively(header)
	header.levelPrint()
	fmt.Print("(:")

	// Output:
	// 8 8 7 9 2 4 7 mirror:8 7 8 2 9 7 4 (:
}

func ExampleMirrorRecursively() {
	header := initTree1()
	header.levelPrint()

	fmt.Print("mirror:")
	MirrorRecursively(header)
	header.levelPrint()
	fmt.Print("(:")

	// Output:
	// 8 8 7 9 2 4 7 mirror:8 7 8 2 9 7 4 (:
}
