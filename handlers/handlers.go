package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	giveCookie(rw, r)
	readHTML(rw, r)
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
		cookie, err := r.Cookie("Test")
		fmt.Fprintf(rw, "Hello there! Here's a value from cookie: %s", cookie.Value)

		if err != nil {
			fmt.Fprintf(rw, "Error: no cookie to read from.")
		}

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

func Params(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "params: %s \n", r.URL.Path)
}

func readHTML(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./html/index.html")
}

func giveCookie(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "Test",
		Value:    "Some value",
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(rw, &cookie)
}

func getValues() ([]byte, error) {
	data := map[string]int{
		"Temperature": rand.Intn(30),
		"Pressure":    900 + rand.Intn(130),
		"Humidity":    20 + rand.Intn(60),
	}

	return json.Marshal(data)
}

func PrintHeaders(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("request headers: ")
	for i, j := range r.Header {
		fmt.Println(i, j)
	}
}
