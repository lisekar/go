package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func chainMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// logging middleware example
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// authentication middleware example
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dummy authentication check
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	fmt.Println("Middleware chaining example.")
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(helloHandler)
	// Chain middlewares
	chain := chainMiddlewares(finalHandler, loggingMiddleware, authMiddleware)
	mux.Handle("/hello", chain)
	fmt.Println("Starting server at port 9100")
	http.ListenAndServe(":9100", mux)
}
