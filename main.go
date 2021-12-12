package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo-address", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, r.RemoteAddr)
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
