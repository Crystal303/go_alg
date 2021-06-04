package stack

import "sync"

type Node struct {
	next  *Node
	value int
}

type Stack struct {
	top *Node
	len int
	mu  sync.RWMutex
}

func NewStack(data ...int) *Stack {
	s := &Stack{
		top: nil,
		len: 0,
		mu:  sync.RWMutex{},
	}
	s.len = len(data)
	for _, v := range data {
		n := &Node{
			next:  nil,
			value: v,
		}
		n.next = s.top
		s.top = n
	}
	return s
}

func (s *Stack) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.len == 0
}

func (s *Stack) Peek() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.top.value
}

func (s *Stack) Pop() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.top != nil {
		temp := s.top
		s.top = s.top.next
		s.len--
		return temp.value
	}
	return 0
}

func (s *Stack) Push(elem int) {
	node := &Node{
		next:  nil,
		value: elem,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	node.next = s.top
	s.top = node
	s.len++
}
