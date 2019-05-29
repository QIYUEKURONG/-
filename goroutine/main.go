package main

import (
	"fmt"
	"time"
)

/*
协程(轻量级的线程)
1：非抢占式多任务处理，由协程主动交出控制权
2：编译器/解释器/虚拟机层面的多任务。
3：多个协程可能在一个或多个线程上面运行
4：任何函数只需要加上go就能送给调度器运行
   不需要再定义时区分是否是异步函数
    调度器再合适的点进行切换
	go run -race main.go 去检测冲突
切换的点都有哪些：
I/O，select
channel
等待锁
函数调用(有时)
runtime.Gosched()
只是参考
*/

func main() {
	/* for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("Hello world\n", i)  //进行了I/O操作会主动交出控制权
		}(i)

	} */
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			a[i]++
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
