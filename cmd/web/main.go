package main

import (
	"fmt"
	"net/http"
	"todo-app/config"
	"todo-app/internals/handlers"
)

func main() {

	var cfg config.Config
	cfg.Port = ":3000"

	http.HandleFunc("/api", handlers.Api)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Printf("server starting on : http://localhost%s/ !", cfg.Port)

	http.ListenAndServe(cfg.Port, nil)

}
