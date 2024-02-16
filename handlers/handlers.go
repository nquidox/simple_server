package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	v := 0

	if v != 0 {
		fmt.Println("request payload: \n")
		fmt.Println("request method: ", r.Method)
		fmt.Println("request headers: ")
		for i, j := range r.Header {
			fmt.Println(i, j)
		}
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(rw, "Hello there! \n")

	case http.MethodPost:
		s, _ := io.ReadAll(r.Body)
		fmt.Fprintf(rw, "POST-ed: %s \n", s)

	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TempHandler(rw http.ResponseWriter, r *http.Request) {
	temp := rand.Intn(30)
	fmt.Fprintf(rw, "Random temp goes here: %d \n", temp)
}

func JsonValues(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, err := getValues()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(rw, string(response))
}

func MuxTest(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Mux Test \n")
}

func getValues() ([]byte, error) {
	data := map[string]int{
		"Temperature": rand.Intn(30),
		"Pressure":    900 + rand.Intn(130),
		"Humidity":    20 + rand.Intn(60),
	}

	return json.Marshal(data)
}
