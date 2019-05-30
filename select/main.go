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

func main() {
	//DoWorker()
	var c1, c2 = GetValue(), GetValue()
	c := CreateSimpWorker()
	for {
		select {
		case n := <-c1:
			fmt.Println("This is c1")
			c <- n
			//fmt.Println("Receive from c1", n)
		case n := <-c2:
			fmt.Println("This is c2")
			c <- n
			//fmt.Println("Receive from c2", n)
		}
	}
}
