package main

import (
	"fmt"

	"github.com/QIYUEKURONG/tutorials/interface/china"
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
func main() {
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
	var g GetPersonAndSex
	g = china.China{"juwenjie", "nan"}
	GetAll(g)
}

/*
1:三个有用的接口
writer
reader
Stringer







*/
