package main

import (
	"fmt"
	"log"
	"net/http"
	hnd "simple_server/handlers"
	mw "simple_server/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/mux", hnd.MuxTest)
	mux.HandleFunc("/", hnd.RootHandler)
	mux.HandleFunc("/temp", hnd.TempHandler)
	mux.HandleFunc("/json", hnd.JsonValues)
	mux.HandleFunc("/middleware", mw.TestMiddleware(hnd.PrintHeaders))

	paramsPath := "/params/"
	mux.Handle(paramsPath, http.StripPrefix(paramsPath, http.HandlerFunc(hnd.Params)))

	fmt.Println("Simple HTTP server")
	log.Fatal(http.ListenAndServe(":9001", mux))
}
