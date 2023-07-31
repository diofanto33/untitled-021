package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diofanto33/web-pass-generator/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.GeneratorHandler)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
