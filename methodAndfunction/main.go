package main

import "fmt"

// adder is a function
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// Method is a method
type Method struct {
	name string
	num  int
}

//PrintfNameAndNum function can peintf name and num . m is a recvier(接收者)
func (m Method) PrintfNameAndNum() {
	fmt.Printf("printf all=%s name and num=%d", m.name, m.num)
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", a(i))
	}
}

/*
	接收者
	1：接收者的命名规定。
    社区约定的接收器命名是类型的一个或两个字母的缩写(像 c 或者 cl 对于 Client)。
	2：接收者不能是指针。
	type Q *MyInt
	func (q Q)Print(){fmt.Println(q)}  x.invalid receiver type Q (Q is a pointer type)




	函数式编程(Go语言)
	函数是一等公民：参数，变量，返回值都可以是函数
	"正统" 函数式编程
	不可变性：不能有状态，只有常量和函数
	函数只能与一个参数
	闭包
*/
