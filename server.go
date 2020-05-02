package main

import (
	"fmt"
	"net/http"
	"./controllers"
)

func main() {
	http.HandleFunc("/", controllers.InfoHandler)
	http.HandleFunc("/submit", controllers.CodeHandler)

	fmt.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}