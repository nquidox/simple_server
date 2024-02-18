package middleware

import (
	"fmt"
	"net/http"
)

func TestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Message from middleware function. \n")
		next(rw, r)
	}
}
