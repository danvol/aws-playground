package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service 2 request path, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Web Service 2 starts listening at:" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
