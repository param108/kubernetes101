package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html>ok</html>"))
}
func main() {
	server := &http.Server{
		Addr: "0.0.0.0:8080",
	}

	handler := mux.NewRouter()
	handler.HandleFunc("/ping", pingHandler)

	server.Handler = handler

	server.ListenAndServe()
}
