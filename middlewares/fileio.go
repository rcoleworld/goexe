package middlewares

import (
	"fmt"
	"os"
)

// FileCreator ...
func FileCreator(body string) {
	file, err := os.Create("/Users/reginaldthomas/Development/PersonalProjects/goexe/codeexecution/tests/Test.py")
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