package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename = "aa.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
