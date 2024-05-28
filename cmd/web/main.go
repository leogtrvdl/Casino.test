package main

import (
	"fmt"
	"net/http"

	"Casinotest/config"
	"Casinotest/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/blackjack", handlers.Blackjack)
	http.HandleFunc("/roulette", handlers.Roulette)

	fmt.Println("(http://localhost:8080) - Server started on port", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
