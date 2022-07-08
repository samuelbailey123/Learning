package main

import (
	"fmt"
	"log"
	"net/http"
	"playground/sites"
)

func main() {
	http.HandleFunc("/", sites.SourceCheck)
	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
