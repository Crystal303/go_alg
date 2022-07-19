package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieNode_Search(t1 *testing.T) {
	root := NewTrieNode(true)
	root.Insert("你好吗", 0)
	root.Insert("你好吗", 1)
	root.Insert("你好吗", 2)
	root.Insert("你好吗", 3)
	root.Insert("你好吗", 4)
	root.Insert("暗号", 5)
	root.Insert("轨迹", 6)
	root.Insert("红模仿", 7)
	root.Insert("你好", 8)

	assert.Equal(t1, 5, root.Search("你好吗"))
	assert.Equal(t1, 6, root.SearchPrefix("你好"))
	assert.Equal(t1, 0, root.SearchFirstOccur("你好吗"))
	assert.Equal(t1, -1, root.SearchFirstOccur("你不在"))
	assert.Equal(t1, 5, root.SearchFirstOccur("暗号"))
}
