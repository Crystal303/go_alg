package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	var a, b []string = nil, []string{}
	assert.False(t, reflect.DeepEqual(a, b)) // false
	assert.True(t, Equal(a, b))              // true

	var c, d map[string]int = nil, make(map[string]int)
	assert.False(t, reflect.DeepEqual(c, d)) // false
	assert.True(t, Equal(c, d))              // true

	type link struct {
		value string
		tail  *link
	}
	l1, l2, l3 := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}

	l1.tail, l2.tail, l3.tail = l2, l1, l3
	assert.True(t, Equal(l1, l1))
	assert.True(t, Equal(l2, l2))
	assert.True(t, Equal(l3, l3))
	assert.False(t, Equal(l2, l1))
	assert.False(t, Equal(l3, l1))
}
