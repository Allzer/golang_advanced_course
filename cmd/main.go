package main

import (
	"fmt"
	"http-server/configs"
	"http-server/internal/pingpong"
	"net/http"
)

func main() {

	config := configs.LoadConfig()

	router := http.NewServeMux()
	pingpong.NewPingPongHandler(router)

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Сервер хапущен на 5000 порте")
	server.ListenAndServe()
}
