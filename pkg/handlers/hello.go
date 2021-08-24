package handlers

import (
	"fmt"
	"log"
	"net/http"
)

const c1 = "Hello"

type Hello struct {
	l *log.Logger
}

//initialises the struct.
func NewHello(l *log.Logger) *Hello {

	return &Hello{l}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s", []byte("Harshal"+c1))
}
