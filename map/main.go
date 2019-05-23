package main

import "fmt"

func maxLengthOfSubstr(str string) int {

	lastOccurrend := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(str) {

		size, ok := lastOccurrend[ch]
		if ok && size >= start {
			start = lastOccurrend[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurrend[ch] = i
	}

	return maxLength
}

func main() {
	length := maxLengthOfSubstr("abcas")
	fmt.Println(length)

}

//1.map是hash map。底层是无序的。
//2：map除了slice，map，function的内建类型都可以作为key
//   当然我们自定义的Struct也可以当作key
