package stack

// push 为压栈序列，pop 为弹出序列
// [1 2 3 4 5] [4 5 3 2 1]
func InPopOrder(push, pop []int) (ret bool) {
	if push != nil && pop != nil && len(push) == len(pop) {
		stack := NewStack()
		for len(push) > 0 {
			if pop[0] != push[0] {
				stack.Push(push[0])
				push = push[1:]
			} else {
				pop = pop[1:]
				push = push[1:]
			}
		}
		for !stack.IsEmpty() && len(pop) != 0 {
			if stack.Peek() == pop[0] {
				stack.Pop()
				pop = pop[1:]
				continue
			}
			break
		}
		if stack.IsEmpty() && len(pop) == 0 {
			ret = true
		}
	}
	return ret
}
