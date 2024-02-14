package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello!")
}

func TempHandler(rw http.ResponseWriter, r *http.Request) {
	temp := rand.Intn(30)
	fmt.Fprintf(rw, "Random temp goes here: %d \n", temp)
}

func JsonValues(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, err := getValues()
	fmt.Println("JsonValues: ", string(response))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(rw, string(response))
}

func getValues() ([]byte, error) {
	data := map[string]int{
		"Temperature": rand.Intn(30),
		"Pressure":    900 + rand.Intn(130),
		"Humidity":    20 + rand.Intn(60),
	}

	return json.Marshal(data)
}
