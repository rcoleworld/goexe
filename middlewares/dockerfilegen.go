package middlewares


import (
	"fmt"
	"os"
	"../keys"
)


func DockerFileGen(taskID string) {
	dockerfile := fmt.Sprintf("FROM python\nRUN mkdir /home/%s\nCOPY %s.py /home/%s/%s.py\nCMD python /home/%s/%s.py", taskID, taskID, taskID, taskID, taskID,taskID)
	filePath := keys.FilePath + taskID + "/" + "Dockerfile"
	file, err := os.Create(filePath)
	
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(dockerfile)
		
	if err != nil {
		panic(err)
	}
	file.Sync()
}