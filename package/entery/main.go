package main

import "fmt"

// Server ...
type Server struct {
	ServerIpandPort string
}

// Start ...
func (s *Server) Start() {
	fmt.Println("Start to a Server ", s.ServerIpandPort)
}

// Client ...
type Client struct {
	Server
	// Ip ...
	IP string
	//Port ...
	Port uint32
}

// Start ...
func (c *Client) Start() {
	fmt.Println("Start to a client ", c.IP, " ", c.Port)
}

func main() {
	/* p := arr.Arr{}
	p.Age = 10
	p.Name = "juwenjie"
	fmt.Println(p) */
	/* 	q := arr.Queue{}
	   	q.QueuePush(1)
	   	q.QueuePush(2)
	   	q.QueuePush(3)
	   	q.ShowQueue()
	   	q.QueuePop()
	   	q.QueuePop()
	   	q.ShowQueue()
	   	fmt.Println(q.IsEmpth())
	   	q.QueuePop()
	   	fmt.Println(q.IsEmpth()) */
	client := Client{IP: "127.0.0.1", Port: 8080}
	client.Server = Server{"127.0.0.1:9000"}
	client.Start()

}

/* 注意
1：包名不一定和文件夹的名字是相同的
2：main包里面有可执行的main程序
3：如何扩展系统类型和别人的类型
  定义别名和使用组合
4: 在Go语言里面体现的继承就是，当嵌套类有某一个方法的时候，包含类可以直接调用。体现的覆盖就是，当嵌套类和包含类的方法是一样的话，那么调用的就是包含类的。
5：go get获取第三方库
6：使用gopm来获取无法下载的包
*/
