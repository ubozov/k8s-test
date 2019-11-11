package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q, %q", html.EscapeString(r.URL.Path), time.Now())
	})

	log.Println("listening port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
