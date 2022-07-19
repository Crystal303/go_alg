package linkedlist

//- 根结点为不包含字符，除根结点外每一个节点包含一个字符；
//- 每个节点的所有子节点包含的字符各不相同；
//- 从根节点到某一节点，路径上经过的字符连接起来，为该节点对于的字符串；

type Trier interface {
	// Search 单词数量查询
	Search(string) int
	// SearchPrefix 单词前缀数量查询
	SearchPrefix(string) int
	// SearchFirstOccur 单词首次出现位置
	SearchFirstOccur(string) int
	Insert(string, int)
}

// TrieNode 字典树节点
type TrieNode struct {
	IsRoot     bool
	Prefix     int
	Count      int
	FirstOccur int
	TrieNode   map[rune]*TrieNode
}

func (t *TrieNode) Search(s string) int {
	if !t.IsRoot {
		return -1
	}

	current := t
	for _, v := range s {
		node, ok := current.TrieNode[v]
		if !ok {
			return -1
		}
		current = node
	}
	return current.Count
}

func (t *TrieNode) SearchPrefix(s string) int {
	if !t.IsRoot {
		return -1
	}

	current := t
	for _, v := range s {
		node, ok := current.TrieNode[v]
		if !ok {
			return -1
		}
		current = node
	}
	return current.Prefix
}

func (t *TrieNode) SearchFirstOccur(s string) int {
	if !t.IsRoot {
		return -1
	}

	current := t
	for _, v := range s {
		node, ok := current.TrieNode[v]
		if !ok {
			return -1
		}
		current = node
	}
	return current.FirstOccur
}

func (t *TrieNode) Insert(s string, idx int) {
	if !t.IsRoot {
		return
	}

	current := t
	for _, v := range s {
		node, ok := current.TrieNode[v]
		if !ok {
			node = NewTrieNode(false)
			current.TrieNode[v] = node
		}
		node.Prefix++
		current = node
	}

	if current.FirstOccur == -1 {
		current.FirstOccur = idx
	}
	current.Count++
}

func NewTrieNode(isRoot bool) *TrieNode {
	return &TrieNode{
		IsRoot:     isRoot,
		Prefix:     0,
		Count:      0,
		FirstOccur: -1,
		TrieNode:   make(map[rune]*TrieNode),
	}
}
