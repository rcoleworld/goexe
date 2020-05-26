package middlewares

import (
	"fmt"
	"os"
	"../keys" // create a keys file to store your file path
)

// Creates a file to execute
func FileCreator(body, fileName, fileExtension, dir string) {
	folderPath := keys.FilePath + dir
	err := os.Mkdir(folderPath, 0755)
	filePath := keys.FilePath + dir + "/" + fileName + "." + fileExtension
	file, err := os.Create(filePath) // file path stored in keys
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