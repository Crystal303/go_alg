package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInPopOrder(t *testing.T) {
	push := []int{1, 2, 3, 4, 5}

	pop := []int{4, 5, 3, 2, 1}
	ret := InPopOrder(push, pop)
	assert.True(t, ret)

	pop = []int{4, 3, 5, 2, 1}
	ret = InPopOrder(push, pop)
	assert.False(t, ret)
}
