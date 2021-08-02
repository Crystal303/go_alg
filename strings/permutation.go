package strings

import (
	"fmt"
)

// 打印字符串的所有排列
func Permutation(str string) {
	permutation([]byte(str), 0)
}

func permutation(str []byte, index int) {
	l := len(str)
	if l == index+1 {
		//fmt.Printf("%s\n", string(str))
		return
	}

	for i := index; i < l; i++ {
		data := make([]byte, l)
		copy(data, str)

		temp := data[i]
		data[i] = data[index]
		data[index] = temp

		permutation(data, index+1)
	}
}

// 打印字符串的所有组合
// n 个字符中 m 长度的组合
// 1个字符和 n-1 个字符
// 剩余的 n-1 个字符选取 m 个字符
// 剩余 n-1 个字符选取 m-1 个
func Combination(str string) {
	l := len(str)
	// n 个字符 有1-n种情况
	for i := 1; i <= l; i++ {
		combination(str, "", i)
	}
}

func combination(str, curStr string, count int) {
	if count == 0 {
		fmt.Println(curStr)
		return
	}

	curStr += str[:1]
	count--
	combination(str[1:], curStr, count)

	if count < len(str)-1 {
		combination(str[1:], curStr[:len(curStr)-1], count+1)
	}
}
