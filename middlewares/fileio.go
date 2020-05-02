package middlewares

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileCreator ...
func FileCreator(body string) {
	absPath, _ := filepath.Abs("../codeexecution/tests/Test.py")
	file, err := os.Create(absPath)
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