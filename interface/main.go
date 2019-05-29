package main

import (
	"fmt"

	"github.com/QIYUEKURONG/tutorials/interface/china"
	"github.com/QIYUEKURONG/tutorials/interface/food"
	"github.com/QIYUEKURONG/tutorials/interface/japan"
)

// Person interface can get name
type Person interface {
	GetName()
}

// Sex interface can get sex
type Sex interface {
	GetSex()
}

// GetPerson ...
func GetPerson(p Person) {
	p.GetName()
}

// GetSex ...
func GetSex(s Sex) {
	s.GetSex()
}

// GetPersonAndSex ....
type GetPersonAndSex interface {
	Person
	Sex
}

// GetAll ...
func GetAll(g GetPersonAndSex) {
	g.GetName()
	g.GetSex()
}

// printfType ...
func printfType(p Person) {
	switch p.(type) {
	case china.China:
		fmt.Println("this is china type")
	case japan.Japan:
		fmt.Println("this is japan type")
	}
}

// Food interface can record a new food
type Food interface {
	DescribeFood()
	PrintfFood()
}

// World interface can record a new place
type World interface {
	Shanxi
}

// Shanxi is a place
type Shanxi interface {
	feature()
}

func main() {

	var f Food
	f = &food.Banana{Name: "banana"}
	f.DescribeFood()
	f.PrintfFood()
	/* var p Person
	p = china.China{"juwenjie"}
	printfType(p)
	p.GetName() */
	/* t, ok := p.(china.China)
	if ok {
		fmt.Println("the type is:", t)
	} else {
		fmt.Println("the type is not:", t)
	} */
	/* 	fmt.Printf("!!!%T %v\n", p, p)
	   	p = japan.Japan{"canglaoshi"}
	   	printfType(p)
	   	p.GetName() */
	/* 	q := queue.Queue{}
	   	q.QueuePush(1)
	   	q.QueuePush(2)
	   	q.QueuePush(3)
	   	q.QueuePush("juwenjie")
	   	q.QueuePush("liujiali")
	   	q.QueuePop()
	   	q.QueuePop()
	   	q.QueuePop()
	   	q.QueuePop()

	   	q.Show() */
	//var g GetPersonAndSex
	//g = china.China{"juwenjie", "nan"}
	//GetAll(g)

}

/*
1:三个有用的接口
writer
reader
Stringer

1:接口里面的方法，在实现它的结构体里面必须要全部实现
2：接口里面可以镶嵌入新的接口。但是不能是它本身和结构体
3:接口类型实现了普通类型转换的两种形式：
   Comma-ok断言和Type-switch测试。
4：Go语言动态类型的实现通常需要编译器静态检查的支持。




*/
