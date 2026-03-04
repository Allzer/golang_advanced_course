package main

import (
	"fmt"
	"net/http"
	"http-server/internal/pingpong"
)

func main() {

	router := http.NewServeMux()
	pingpong.NewPingPongHandler(router)

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Сервер хапущен на 5000 порте")
	server.ListenAndServe()
}
