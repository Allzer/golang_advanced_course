package pingpong

import (
	"fmt"
	"net/http"
)

type PingPongHandler struct {}

func NewPingPongHandler(router *http.ServeMux) {
	handler := &PingPongHandler{}
	router.HandleFunc("/ping", handler.PingPong())
}

func (handler *PingPongHandler) PingPong() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Pong")
	}
}