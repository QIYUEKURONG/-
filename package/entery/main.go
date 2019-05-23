package main

import (
	"fmt"

	arr "github.com/QIYUEKURONG/tutorials/package/entery/arrary"
)

func main() {
	p := arr.Arr{}
	p.Age = 10
	p.Name = "juwenjie"
	fmt.Println(p)
}

// 注意
//1：包名不一定和文件夹的名字是相同的
//2：main包里面有可执行的main程序
