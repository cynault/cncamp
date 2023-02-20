package main

import (
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	header := r.Header
	for k, v := range header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(100)
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	println(r.RemoteAddr, " statusCode : ", 100)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":80", nil)
}
