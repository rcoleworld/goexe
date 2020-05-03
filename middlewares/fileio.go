package middlewares

import (
	"fmt"
	"os"
	"../keys" // create a keys file to store your file path
)

// FileCreator ...
func FileCreator(body string) {
	file, err := os.Create(keys.FilePath) // file path stored in keys
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n, err := file.WriteString(string(body))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	file.Sync()
}