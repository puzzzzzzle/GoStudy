package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = "D:/Go/bin/go"
		script := "D:/workspace/goPro/" + r.URL.Path
		log.Println(handler.Path)
		handler.Dir = "D:/workspace/goPro/"
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.Env = append(handler.Env, "GOPATH=D:/Go/gopath")
		handler.Env = append(handler.Env, "GOROOT=D:/Go")
		log.Println(handler.Args)

		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}