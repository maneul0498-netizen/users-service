package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func health(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		//Message: "service running, testing github web hook 2",
		Message: Message("Manuel"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Message(m string) string {
	return fmt.Sprintf("Hello %s !!!", m)
}

func main() {
	http.HandleFunc("/health", health)

	log.Println("server running on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
