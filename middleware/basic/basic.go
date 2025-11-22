package main

import (
	"fmt"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Println("Request processed in:", duration)
	})
}
func helloHAndler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	fmt.Println("This is a placeholder file.")
	mux := http.NewServeMux()
	// mux.HandleFunc("/hello", helloHAndler)
	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHAndler)))
	fmt.Println("Starting server at port 9000")
	http.ListenAndServe(":9000", mux)
}
