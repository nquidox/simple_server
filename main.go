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
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/temp", handlers.TempHandler)
	http.HandleFunc("/json", handlers.JsonValues)

	fmt.Println("Simple HTTP server")
	log.Fatal(http.ListenAndServe(":9001", mux))

}
