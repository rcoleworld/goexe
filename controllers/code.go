package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

// CodeHandler ...recieves code sent by the client.
func CodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body.", http.StatusInternalServerError)
		}
		fmt.Fprint(w, strconv.Quote(string(body)))
	} else {
		http.Error(w, "Request type not allowed", http.StatusMethodNotAllowed)
	}
}