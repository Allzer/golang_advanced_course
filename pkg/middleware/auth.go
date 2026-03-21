package middleware

import (
	"context"
	"fmt"
	"http-server/configs"
	"http-server/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJwt(config.Auth.Secret).Parse(token)
		fmt.Println(isValid)
		fmt.Println(data)

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)


		next.ServeHTTP(w, req)
	})
}