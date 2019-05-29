package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/go/", func(writer http.ResponseWriter,
		request *http.Request) {
		path := request.URL.Path[len("/go/"):]
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(all)
		
	})
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic("http lister and server find error")
	}
	
}
