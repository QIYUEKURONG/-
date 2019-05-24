package queue

import "fmt"

// Queue ...
type Queue []interface{}

// QueuePush ...
func (q *Queue) QueuePush(value interface{}) {
	*q = append(*q, value)
}

//假如我们要限制push或pop的值
//func (Q *Queue) QueuePush(value int) {}  或者func (q *Queue) QueuePush(value interface{}) {*q=append(*q, value.(int))}  后面这个是在运行的时候才能直到的错误。

// QueuePop ...
func (q *Queue) QueuePop() interface{} {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

// Show ...
func (q *Queue) Show() {
	fmt.Println(*q)
}
