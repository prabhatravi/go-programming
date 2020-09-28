package main

import (
	"fmt"
	"net/http"
	"os"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get hostname")
	}
	// create response binary data
	data := append([]byte("Hello "), hostname...) // slice of bytes
	// write `data` to response
	res.Write(data)
}

func main() {

	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":9000", handler)

}
