package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"../middlewares"
	"../codeexecution"
)

func CodeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body.", http.StatusInternalServerError)
		}
		fmt.Println(strconv.Quote(string(body)))
		// creates a file from the response
		fileExtension := r.Header.Get("File-Ext")
		taskID := r.Header.Get("Task-ID")
		language := r.Header.Get("Language")	
		middlewares.FileCreator(string(body), taskID, fileExtension, taskID)
		middlewares.DockerFileGen(taskID, language)
		fmt.Fprint(w, codeexecution.DockerRun(taskID))
	} else {
		http.Error(w, "Request type not allowed", http.StatusMethodNotAllowed)
	}
}
