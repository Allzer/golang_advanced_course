package main

import (
	"fmt"
	"http-server/configs"
	"http-server/internal/auth"
	"http-server/internal/link"
	"http-server/pkg/db"
	"http-server/pkg/middleware"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()

	db := db.NewDb(conf)

	router := http.NewServeMux()

	//reposititories

	linkRepository := link.NewLinkRepository(db)

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	//connection
	server := http.Server{
		Addr:    ":5000",
		Handler: stack(router),
	}

	fmt.Println("Сервер хапущен на 5000 порте")
	server.ListenAndServe()
}
