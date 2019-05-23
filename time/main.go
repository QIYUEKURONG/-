package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	current, mon, day := time.Now().Date()
	fmt.Printf("%d-%d-%d", current, mon, day)
	Nanosecond := time.Now().Nanosecond()
	fmt.Printf("-%d ", Nanosecond)
	var conn net.Conn
	conn.writeBuf()
}

//跳出for循环。说明服务器有这个数据
/* fmt.Println("the file exits", *Response)
//建一个Piece结构体 要开始接收数据了
//piecebuff := make([]byte, 5048)
piecereq.NewObject(task, piecereq)
sedbuff, err3 := piecereq.EncodeBody(piecereq)
if err3 != nil {
	fmt.Println("Sorry ! the piece encode error")
	os.Exit(-1)
}
conn.Write(sedbuff)
//接收数据
piecebuff := make([]byte, 18000)
fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!")
count, _ := conn.Read(piecebuff)
fmt.Println("count", count)
fmt.Println("piece buff", piecebuff)
fmt.Println(len(piecebuff))
*/
/* pieceres := new(protocol.PieceResponse)
*pieceres, err = pieceres.DecodeBody(piecebuff)
fmt.Printf("%s", pieceres.PieceData) */

//转换成二进制 然后发送给服务器
/* 	var message protocol.BlockRequest
AssignmentStruct(&message, task)
buf, err := SerializateBinary(&message)
if err != nil {
	fmt.Println("sorry! Object convert binary error")
	os.Exit(-1)
}
fmt.Println(buf)
//buff := make([]byte, 0)
//fmt.Println(buff)
// 循环去接收服务器发送来的数据
buff := make([]byte, 0)
_, err1 := conn.Read(buff)
if err1 != nil {
	fmt.Println("client read error")
	os.Exit(-1)
}
data, err1 := UnserializateBinary(buff)
if err1 != nil {
	fmt.Println("sorry! UnserializateBinary  error")
}
//写入文件里面
file.Write(([]byte)(data.FileIndex))
//如果文件的尺寸不是0的话，就继续请求文件
if data.FileSize != 0 {
	task.FileIndex = data.FileIndex
	task.BlockIndex = (uint)(data.FileOffset)
} else {
	fmt.Println("file download sucess")
	break
}
*/
