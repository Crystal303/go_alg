package linkedlist

import "math/rand"

const (
	LevelMax = 4
	P        = 0.5
)

type SkipLister interface {
	// Find 查找
	Find(int) bool
	// FindRange 范围查询
	FindRange(int, int) []int
	Insert(int)
	Delete(int) bool
	Print()
}

// SkipNode 跳表节点
type SkipNode struct {
	Value    int
	SkipNode []*SkipNode
}

func NewSkipNode(value, level int) *SkipNode {
	node := &SkipNode{
		Value: value,
	}
	node.initLevel(level)
	return node
}

func (s *SkipNode) initLevel(level int) {
	if level < 0 {
		return
	}
	s.SkipNode = make([]*SkipNode, level)
	for i := range s.SkipNode {
		s.SkipNode[i] = new(SkipNode)
	}
}

func (s *SkipNode) GetLevel() int {
	return len(s.SkipNode)
}

func (s *SkipNode) Find(value int) bool {
	level := s.GetLevel()
	current := s
	if current.Value == value {
		return true
	}

	for i := level - 1; 0 <= i; i-- {
		for 0 < current.GetLevel() {
			if value == current.SkipNode[i].Value {
				return true
			}
			if value < current.SkipNode[i].Value {
				break
			}
			// 索引不在最后一层
			if current.SkipNode[i].SkipNode[i].GetLevel() == 0 {
				break
			}
			current = current.SkipNode[i]
		}
	}
	return false
}

func (s *SkipNode) FindRange(i int, i2 int) []int {
	//TODO implement me
	panic("implement me")
}

func (s *SkipNode) Insert(value int) {
	level := RandomLevel(LevelMax, P)

	node := NewSkipNode(value, level)
	current := s
	for i := level - 1; 0 <= i; i-- {
		for 0 < current.SkipNode[i].GetLevel() {
			if v1 := current.SkipNode[i].Value; value < v1 {
				break
			}
			current = current.SkipNode[i]
		}
		node.SkipNode[i] = current.SkipNode[i]
		current.SkipNode[i] = node
	}
}

func (s *SkipNode) Delete(i int) bool {
	// 删除第一个元素和最后一个元素
	//TODO implement me
	panic("implement me")
}

func RandomLevel(maxLevel int, p float64) int {
	level := 1
	if maxLevel < 0 ||
		p < 0 ||
		1 < p {
		return level
	}
	n := int(100 * p)
	for i := 0; i < maxLevel; i++ {
		if rand.Intn(100) < n {
			level++
			continue
		}
		return level
	}

	return level
}

func (s *SkipNode) Print() {
	level := s.GetLevel()

	for i := level - 1; 0 <= i; i-- {
		print("head: ")

		current := s
		for 0 < current.GetLevel() {
			print(current.Value, " ")
			current = current.SkipNode[i]
		}
		print("\n")
	}
}
