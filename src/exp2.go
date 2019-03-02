package main

import (
	"io"
	"log"
	"net/http"
)

func goServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>This is a go web server</h1><img src=\"https://golang.google.cn/doc/gopher/frontpage.png\"/>")
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/go", logPanics(goServer))
	log.Println("Hosting server on 0.0.0.0:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
