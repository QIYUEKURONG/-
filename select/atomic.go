package main

import (
	"fmt"
	"sync"
	"time"
)

type atomic struct {
	value int
	lock  sync.Mutex
}

func (a *atomic) AddAtomic() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}
func (a *atomic) GetValue() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	/* var a atomic
	a.AddAtomic()
	go func() {
		a.AddAtomic()
	}()
	time.Sleep(time.Microsecond)
	fmt.Println(a.GetValue()) */
	var a atomic
	a.AddAtomic()
	go func() {
		a.AddAtomic()
	}()
	time.Sleep(time.Microsecond)
	fmt.Println(a.GetValue())
}
