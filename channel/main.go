package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker ...
func Worker() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("goroutine = %d and chan value = %d\n", i, <-c)
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	time.Sleep(time.Millisecond)
}

// ChanValue ...
func ChanValue() chan int {
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("chan value is = %d\n", <-c)
		}
	}()
	return c
}

// BufferChan ...
func BufferChan() chan int {
	c := make(chan int, 5)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		c <- 5
		close(c)
	}()

	return c
}

// RealWorker can record bool type's printf finish and printf channl
type RealWorker struct {
	c  chan int
	wg *sync.WaitGroup
}

// ToWorker can do work
func ToWorker(r RealWorker) {
	for v := range r.c {
		fmt.Printf("the chan value is: %c\n", v)
		r.wg.Done()
	}
}

// CreateWorker function can create a chan
func CreateWorker(g *sync.WaitGroup) RealWorker {
	c := RealWorker{
		c:  make(chan int),
		wg: g,
	}
	go ToWorker(c)
	return c
}

// WriteChan function can write in chan
func WriteChan() {
	var wg sync.WaitGroup

	var chanbuff [10]RealWorker
	for i := 0; i < 10; i++ {
		chanbuff[i] = CreateWorker(&wg)
	}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		chanbuff[i].c <- 'a' + i
		//<-chanbuff[i].done
	}
	for i := 0; i < 10; i++ {
		chanbuff[i].c <- 'A' + i
	}
	wg.Wait()
	/* for i := 0; i < 10; i++ {
		chanbuff[i].c <- 'A' + i
		//<-chanbuff[i].done
	}
	for i := 0; i < 10; i++ {
		<-chanbuff[i].done
	} */
}

// BufferChanTest ...
func BufferChanTest() chan<- int {
	c := make(chan int, 5)
	go func() {
		for {
			fmt.Printf("chan value=%d\n", <-c)
		}
	}()
	return c
}
func main() {
	//Worker()
	/* 	c := ChanValue()
	   	c <- 1
	   	c <- 2
	   	c <- 3
	   	c <- 4
	   	time.Sleep(time.Millisecond * 2)
	*/
	/* 	c := BufferChan()

	   	for v := range c {
	   		fmt.Println("chan value", v)
	   	} */
	/* 	c := BufferChanTest()
	   	c <- 1
	   	c <- 2
	   	time.Sleep(time.Microsecond * 3) */
	WriteChan()
}

// 不要通过共享内存来通信，而是通过通信来共享内存
/*
可能出现的问题
1：你要尽量避免让主协程循环的去接收消息，这样的话子协程没有关闭的话，那么主协程就会陷入死锁的状态了。


*/
