package apiserver

import (
	"io"
	"net/http"
)

func (a *APIServer) initRouter() {
	a.router.HandleFunc("/ping", a.HandlePing())
}

func (a *APIServer) HandlePing() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "pong")
	}
}
