package http

import (
	"fmt"
	"log"
	"net/http"
)

const (
	addr = ":8080"
)

// Start ...
func Start() {
	fmt.Println("http listen on ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
