package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()

	var wg sync.WaitGroup

	for i := 0; i<10; i++{
		wg.Add(1)

		go func () {
			getResp()
			wg.Done()
		} ()

	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func getResp() {
	url := "https://google.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}

	fmt.Println("Статус-код:", resp.StatusCode)
}
