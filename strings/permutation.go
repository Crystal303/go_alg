package strings

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
