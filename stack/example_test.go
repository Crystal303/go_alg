package stack

import "fmt"

func ExamplePop() {
	s := NewStack(1, 2, 3, 4)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	// 4
	// 3
	// 2
	// 1
}

func ExamplePush() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	// 3
	// 2
	// 1
}

func ExampleLen() {
	s := NewStack(1, 2, 3)
	fmt.Println(s.Len())
	s.Push(4)
	fmt.Println(s.Len())
	s.Pop()
	fmt.Println(s.Len())

	// Output:
	// 3
	// 4
	// 3
}
