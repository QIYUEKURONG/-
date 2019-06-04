package main

import (
	"fmt"
	"math/rand"

	"time"
)

/*
// Worker ...
type Worker struct {
	w  chan int
	wg *sync.WaitGroup
}

// DoWorker ...
func DoWorker() {
	var wg sync.WaitGroup
	var work [10]Worker
	for i := 0; i < 10; i++ {
		work[i] = CreatrWorker(&wg)
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		work[i].w <- 'a' + i
	}
	wg.Wait()

}

// PrintfChan ...
func PrintfChan(w Worker) {
	for v := range w.w {
		fmt.Printf("chan value: %c\n", v)
		w.wg.Done()
	}
}

// CreatrWorker ...
func CreatrWorker(wg *sync.WaitGroup) Worker {
	c := Worker{
		w:  make(chan int),
		wg: wg,
	}
	go PrintfChan(c)
	return c
} */
// SimpCreateWorker ...
func SimpCreateWorker(c chan int) {
	for v := range c {
		fmt.Printf("chan value: %v\n ", v)
	}

}

// CreateSimpWorker ...
func CreateSimpWorker() chan int {
	c := make(chan int)
	go SimpCreateWorker(c)
	return c
}

// PrintfWorker ...
func PrintfWorker(c chan int) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Println("Printf chan value ", v)
	}
}

// CreateWorker ...
func CreateWorker() chan int {
	c := make(chan int)
	go PrintfWorker(c)
	return c
}

// GetValue ...
func GetValue() chan int {

	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

/*
func main() {
	//DoWorker()
	var c1, c2 = GetValue(), GetValue()
	//c := CreateSimpWorker()
	n := 0
	var w = CreateWorker()
	var valueGroup []int
	tm := time.After(time.Second * 10) //这两个都是定时装置
	tick := time.Tick(time.Second)
	for {

		var active chan int
		var actValue int
		if len(valueGroup) > 0 {
			active = w
			actValue = valueGroup[0]
		}
		select {
		case n = <-c1:
			valueGroup = append(valueGroup, n)
		case n = <-c2:
			valueGroup = append(valueGroup, n)
		case active <- actValue:
			valueGroup = valueGroup[1:]
		case <-time.After(time.Millisecond * 800):
			fmt.Println("time out")
		case <-tick:
			fmt.Println("valueGroup length is:=", len(valueGroup))
		case <-tm:
			fmt.Println("Bye Bye!!!!!")
			return
		}
	}
}*/
