package main

import (
	"fmt"
	"net/http"

	"Casinotest/internal/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/blackjack", handlers.Blackjack)
	http.HandleFunc("/roulette", handlers.Roulette)

	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
