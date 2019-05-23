package main

import "fmt"

type data interface {
	Ok()
}

// Ok ...
func Ok() {
	fmt.Println("ok!!!!")
}

type message interface {
	write()
	read()
}

func main() {
	/* 	u, err := url.Parse("https://www.baidu.com/search?q=dotnet")
	   	if err != nil {
	   		fmt.Println("url parse error")
	   	}
	   	fmt.Printf("url.Scheme : %s\n", u.Scheme)
	   	fmt.Printf("url.Opaque : %s\n", u.Opaque)
	   	fmt.Printf("url.Host : %s\n", u.Host)
	   	fmt.Printf("url.Path: %s\n", u.Path)
	   	fmt.Printf("url.Opaque: %s\n", u.RawPath)
	   	fmt.Printf("url.RawQuery: %s\n", u.RawQuery)
	   	fmt.Printf("url.Fragment: %s\n", u.Fragment) */
	fmt.Println("start")
	var mess message
	val, _ := mess.(data)

	val.Ok()

	fmt.Println("end")
}
