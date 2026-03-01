package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	code := make(chan int)

	var wg sync.WaitGroup

	for i:=0; i<10; i++ {
		wg.Add(1)
		go func() {
			go getResp(code, &wg)
		}()
	}

	go func()  {
		wg.Wait()
		close(code)	
	}()

	for res := range code{
		fmt.Println("Статус-код:", res)
	}
}

func getResp(codeCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://google.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}

	codeCh <- resp.StatusCode
}
