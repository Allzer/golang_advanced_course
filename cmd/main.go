package main

import (
	"fmt"
	"http-server/configs"
	"http-server/internal/auth"
	"http-server/pkg/db"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	
	_ = db.NewDb(conf)

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Сервер хапущен на 5000 порте")
	server.ListenAndServe()
}
