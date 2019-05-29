package main

import "fmt"

/* func tryDefer() {
	for i := 0; i < 30; i++ {
		defer fmt.Printf("%d ", i)
		if i == 30 {
			panic("error!!!")
		}
	}
}

func main() {
	//tryDefer()
	file, err := os.OpenFile("abc.txt", os.O_CREATE, 0666)
	if err != nil {
		value, ok := err.(*os.PathError)
		if ok {
			fmt.Printf("%s %s %v", value.Op, value.Path, value.Err)
		} else {
			panic(err)
		}
	}
	str := "juwenjieshihaoren"
	file.Write([]byte(str))

}
*/
func abnormal() {
	defer func() {
		err := recover()
		err, ok := err.(error)
		if ok {
			fmt.Println("recover error ：", err)
		} else {
			fmt.Println("I don't know this is what error")
		}
	}()
	b := 0
	a := 10 / b
	fmt.Println(a)
}

func main() {
	abnormal()
}

/*
1:defer确保调用在函数结束时发生
2：参数在defer语句时计算(就好像在这一行的时候，进行压栈的处理)
3:defer列表为先进后出
何时去使用：
1：
Open /Close
Lock/Unlock
PrintHeader/PrintFooter
*/
/*
panic的使用
1：停止当前函数的执行
2：一直向上返回，执行每一层的defer
3：如果没有遇见recover，程序退出
*/
/*
	1：仅在defer调用中使用
	2：获取panic的值
	3：如果无法处理，可重新panic
*/
/*
使用：
Go中可以抛出一个panic的异常，然后在defer中通过recover来捕获这个异常，
然后去正常的处理。
*/
