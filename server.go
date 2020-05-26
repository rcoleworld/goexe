package main

import (
	"fmt"
	"net/http"
	"./controllers"
	// "./middlewares"
)

func main() {

	// middlewares.FileCreator("This is a test", "7459af63-3003-4227-be97-48074954aeea", "txt", "7459af63-3003-4227-be97-48074954aeea")
	// middlewares.DockerFileGen("7459af63-3003-4227-be97-48074954aeea")
	http.HandleFunc("/", controllers.InfoHandler)
	http.HandleFunc("/submit", controllers.CodeHandler)
	fmt.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
