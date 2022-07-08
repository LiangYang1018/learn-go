package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hi there")
	http.Get("http://example.com/")
}
