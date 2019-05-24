package arr

import "fmt"

//Queue can record name and age
type Queue []int

// QueuePush ...
func (q *Queue) QueuePush(value int) {
	*q = append(*q, value)
}

// QueuePop ...
func (q *Queue) QueuePop() {
	*q = (*q)[1:]
}

// IsEmpth ...
func (q *Queue) IsEmpth() bool {
	return len(*q) == 0
}

// ShowQueue ...
func (q *Queue) ShowQueue() {
	fmt.Println(*q)
}
