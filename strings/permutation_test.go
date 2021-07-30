package strings

import "testing"

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

func TestSlicePoint(t *testing.T) {
	// 保存的是指针
	a := []byte("ddd")
	// 同一个指针
	b := a
	b[1] = 'c'
	t.Log(string(a), string(b))

	// 不同的指针
	c := make([]byte, len(b))
	copy(c, b)
	c[1] = 'a'
	t.Log(string(c), string(b))
}
