package main

import (
	"fmt"
	"log"
	"net/http"
	"nebula-router-system/src/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}