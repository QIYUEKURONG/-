package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	buff := make([]byte, 10)
	buff[0] = 'a'
	buff[1] = 'b'
	buff[2] = 'c'
	buff[3] = 'd'
	fmt.Println(buff)
	var value uint32
	value = binary.BigEndian.Uint32(buff)
	fmt.Println(value)
}
