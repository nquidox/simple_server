package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello!")
}

func tempHandler(rw http.ResponseWriter, r *http.Request) {
	temp := rand.Intn(30)
	fmt.Fprintf(rw, "Random temp goes here: %d \n", temp)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/temp", tempHandler)

	fmt.Println("Simple HTTP server")
	log.Fatal(http.ListenAndServe("127.0.0.1:9001", nil))

}
