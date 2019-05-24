package main

import (
	"fmt"
	"unicode/utf8"
)

/*  Value ...
func Value(buff []byte) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
	fmt.Printf("%p\n", &buff)
	buff = append(buff, "hao"...)
	fmt.Printf("%s\n", buff)
	fmt.Printf("%p\n", &buff)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
}

func main() {
		buff := make([]byte, 0)
	   	buff = append(buff, "woaini"...)
	   	fmt.Printf("%s\n", buff)
	   	buf := bytes.NewBuffer(buff)
	   	fmt.Printf("%s\n", buf.Bytes())
	   	buff = append(buff, "niaiwoma"...)
	   	fmt.Printf("%s\n", buf.Bytes())
	   	fmt.Printf("%s\n", buff)
	buff := make([]byte, 0)
	fmt.Printf("%p\n", &buff)
	fmt.Printf("buff len=%v and cap=%v\n", len(buff), cap(buff))

	buff = append(buff, "woainihao"...)
	fmt.Printf("buff len=%v and cap=%v\n", len(buff), cap(buff))

	Value(buff)
	fmt.Printf("%s\n", buff)
}*/

/* type Buff struct {
	buf []byte // contents are the bytes buf[off : len(buf)]
	off int    // read at &buf[off], write at &buf[len(buf)]
}

func NewBuff(buf []byte) *Buff { return &Buff{buf: buf} }
func main() {
	 	buff := make([]byte, 0)
	   	buff = append(buff, "aqya"...)
	   	fmt.Println("main len and cap", len(buff), cap(buff))
	   	fun(buff)
	   	fmt.Println("main len and cap", len(buff), cap(buff))
	   	fmt.Printf("%s\n", buff)
	 buff := make([]byte, 10)
	   	fmt.Println("buff len and cap", len(buff), cap(buff))
	   	buf := bytes.NewBuffer(buff)
	   	buf.Write([]byte( "wo  ai"))
	   	fmt.Println("buf len and cap", buf.Len(), buf.Cap())
	   	fmt.Printf("buf %s\n", buf)
}
*/

func main() {
	s := "丽丽是个猪！"
	fmt.Printf("length= %d\n", utf8.RuneCountInString(s))
	bytes := []byte(s)
	ch, size := utf8.DecodeRune(bytes)
	fmt.Printf("%v %d\n", ch, size)
	bytes = bytes[size:]
	fmt.Printf("%s \n", bytes)

}

//rune 相当于Go的char  但是它是int32   byte是uint8rune用来表示Unicode的code point
//使用utf8.RuneCountInString获得字符数量
//使用len获得字节的长度
//使用[]bye获得字节
