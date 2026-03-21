package auth

import (
	"fmt"
	"http-server/configs"
	"http-server/pkg/req"
	"http-server/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

	// router.HandleFunc("GET /auth/login", handler.Login())
	// router.HandleFunc("GET /auth/register", handler.Register())
}

// {
// 	"email": "123@mail.ru",
// 	"password": "123",
// 	"name": "XD",
// }

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		
		userEmail, err := handler.AuthService.Login(body.Email, body.Password)
		fmt.Println(userEmail, err)
		if err != nil {
			return
		}

		fmt.Println(body)

		data := LoginResopnse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			return
		}
		handler.AuthService.Register(body.Email, body.Password, body.Name)
	}
}
