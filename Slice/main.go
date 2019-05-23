package main

import "fmt"

func arrayTest(arr []int) {
	fmt.Println(arr)
	arr[2] = 1000
}
func printfLenandCap(s2 []int) {
	fmt.Printf("len and cap : %d %d\n", len(s2), cap(s2))

}

func main() {

	arrary := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	ar1 := arrary[2:5]
	ar2 := arrary[:]
	ar3 := arrary[3:]
	fmt.Println(ar1, ar2, ar3)
	arrayTest(arrary[:])
	fmt.Println(arrary)

	//delet element form slice
	s1 := make([]int, 0, 16)
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("s1 ", s1)
	//s1 = append(s1[0:1], s1[2:]...)
	//fmt.Println("s1 ", s1)

	//delete front and back
	front := s1[0]
	s1 = s1[1:]
	fmt.Println(s1)
	back := s1[len(s1)-1]
	s1 = s1[0 : len(s1)-1]
	fmt.Println(s1, " ", front, " ", back)

	s2 := make([]int, 10)
	for i := 0; i < 100; i++ {
		printfLenandCap(s2)
		s2 = append(s2, 2*i+1)
	}
}
//1：添加元素如果超越cap，系统就会在底层分配一个更大的底层数组。
//2：如果要删除一个元素，可以使用append这个函数。
