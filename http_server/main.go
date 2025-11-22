package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func getUser1(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello, User1"}
	json.NewEncoder(w).Encode(msg)
}

func getUser2(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello, User2"}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/api/user1", getUser1)
	http.HandleFunc("/api/user2", getUser2)
	fmt.Println("Starting server at port 8000")
	// http.ListenAndServe(":8000", nil)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
