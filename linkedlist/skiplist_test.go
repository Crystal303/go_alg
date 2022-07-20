package linkedlist

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomLevel(t *testing.T) {
	count := 0
	count2 := 0
	times := 10000000
	for i := 0; i < times; i++ {
		level := RandomLevel(4, 0.5)
		if level == 1 {
			count++
		}
		if level == 2 {
			count2++
		}
	}

	assert.Equal(t, "0.50", fmt.Sprintf("%.2f", float64(count)/float64(times)))
	assert.Equal(t, "0.25", fmt.Sprintf("%.2f", float64(count2)/float64(times)))
}

func TestSkipNode_Find(t *testing.T) {
	root := NewSkipNode(math.MinInt, LevelMax)
	root.Insert(2)
	root.Insert(2)
	root.Insert(1)
	root.Insert(5)
	root.Insert(8)
	root.Insert(4)
	root.Insert(3)

	root.Print()

	assert.True(t, root.Find(4))
	assert.True(t, root.Find(8))
	assert.True(t, root.Find(1))
	assert.False(t, root.Find(9))
	assert.False(t, root.Find(0))
	assert.False(t, root.Find(-1))
}
