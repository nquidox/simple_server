package main

import (
	"fmt"
	"log"
	"net/http"
	"simple_server/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/mux", handlers.MuxTest)
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/temp", handlers.TempHandler)
	mux.HandleFunc("/json", handlers.JsonValues)

	fmt.Println("Simple HTTP server")
	log.Fatal(http.ListenAndServe(":9001", mux))

}
