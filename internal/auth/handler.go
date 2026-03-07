package auth

import (
	"encoding/json"
	"fmt"
	"http-server/configs"
	"http-server/pkg/res"
	"net/http"
	"net/mail"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

	router.HandleFunc("GET /auth/login", handler.Login())
	router.HandleFunc("GET /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}

		if payload.Email == "" {
			res.Json(w, "Email is none", 402)
			return
		}

		_, err = mail.ParseAddress(payload.Email)
		if err != nil {
			res.Json(w, "Wrong email", 402)
			return
		}

		if payload.Password == "" {
			res.Json(w, "Password is none", 402)
			return
		}

		data := LoginResopnse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
