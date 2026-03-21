package auth

import (
	"http-server/configs"
	"http-server/pkg/jwt"
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
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		
		userEmail, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		
		token, _ := jwt.NewJwt(handler.Config.Auth.Secret).Create(userEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := LoginResopnse{
			Token: token,
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
