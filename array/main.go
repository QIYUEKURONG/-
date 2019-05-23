package main

import "fmt"

func testIfValue(array [5]int) {
	for i := range array {
		fmt.Println(i)
	}
	array[1] = 10
}
func ifOldValue(arrary, arrary1 [5]int) bool {
	if arrary == arrary1 {
		return true
	}
	return false
}
func main() {
	var arrary [5]int
	arrary1 := [2]int{2, 3}
	arrary2 := [...]int{2, 5, 7, 6, 9}
	fmt.Println(arrary, arrary1, arrary2)
	testIfValue(arrary)
	fmt.Println(arrary)
}

//think:
//1:tne arrary is transport by value
//2:the arrary can to slice
