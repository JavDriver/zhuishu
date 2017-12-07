package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world")
	})

	err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
