package main

import (
	"fmt"
	"log"
	"net/http"
	"simple_server/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/temp", handlers.TempHandler)
	http.HandleFunc("/json", handlers.JsonValues)

	fmt.Println("Simple HTTP server")
	log.Fatal(http.ListenAndServe(":9001", nil))

}
