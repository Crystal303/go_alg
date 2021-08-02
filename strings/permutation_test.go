package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	s := "abc"
	Permutation(s)
}

// go test -v -bench=. -run=BenchmarkPermutation -benchmem
func BenchmarkPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permutation("1234567890")
	}
}

//BenchmarkPermutation
//BenchmarkPermutation-4                 5         202738404 ns/op        99764934 B/op    6235301 allocs/op

func TestCombination(t *testing.T) {
	s := "abc"
	Combination(s)
	//a
	//b
	//c
	//ab
	//ac
	//bc
	//abc
}

// 切片测试，底层数组不变的情况下，slice值改变的影响
func TestSlicePoint(t *testing.T) {
	// 保存的是指针
	a := []byte("ddd")
	// 同一个指针
	b := a
	b[1] = 'c'
	assert.Equal(t, a, b)
	// dcd

	// 不同的指针
	c := make([]byte, len(b))
	copy(c, b)
	c[1] = 'a'
	assert.NotEqual(t, b, c)
	// b: dcd
	// c: dad

	d := make([]byte, len(c))
	copy(d, c)
	sliceIndexChangeData(c)
	assert.NotEqual(t, c, d)
	// c: 2ad
	// d: dad

	e := make([]byte, len(d))
	copy(e, d)
	sliceAppendChangeData(d)
	assert.Equal(t, d, e)
	// dad

	sliceChangePointAppendChangeData(&d)
	assert.Equal(t, d, e)
	// dad

	slicePointAppendChangeData(&d)
	assert.NotEqual(t, d, e)
	// d: ddad
	// e: dad
}

func sliceIndexChangeData(s []byte) {
	if len(s) == 0 {
		return
	}
	s[0] = '2'
}

func sliceAppendChangeData(s []byte) {
	if len(s) == 0 {
		return
	}
	s = append(s, 'd')
}

func slicePointAppendChangeData(s *[]byte) {
	if s == nil ||
		len(*s) == 0 {
		return
	}
	// 改变的是指针指向的地址的值
	*s = append(*s, 'd')
}

func sliceChangePointAppendChangeData(s *[]byte) {
	if s == nil ||
		len(*s) == 0 {
		return
	}
	temp := append(*s, 'd')
	// 换了一个新指针值
	s = &temp
}
