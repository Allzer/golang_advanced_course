package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	path := flag.String("file", "url.txt", "path to URL file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil{
		panic(err.Error())
	}

	urlSlice := strings.Split(string(file), ";")

	recpChan := make(chan int)
	errChan := make(chan error)

	for _, url := range urlSlice{
		go ping(url, recpChan, errChan)
	}
	for range urlSlice {
		select{
		case res := <-recpChan:
			fmt.Println(res)
		case errRes := <- errChan:
			fmt.Println(errRes)
		}
	}
}

func ping(url string, recpChan chan int, errChan chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errChan <- err
		return
	}
	// defer resp.Body.Close()
	recpChan <- resp.StatusCode
}