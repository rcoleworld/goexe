package middlewares


import (
	"fmt"
	"os"
	"../keys"
)

func DockerFileGen(taskID string, lang string) {
	dockerfilePy := fmt.Sprintf("FROM python\nRUN mkdir /home/%s\nCOPY %s.py /home/%s/%s.py\nCMD python /home/%s/%s.py", taskID, taskID, taskID, taskID, taskID,taskID)
	dockerfileGo := fmt.Sprintf("FROM golang:latest\nRUN mkdir /app\nCOPY ./%s.go /app/%s.go \nWORKDIR /app\nRUN go build -o %s .\nCMD [\"/app/%s\"]", taskID, taskID, taskID, taskID, taskID)

	m := map[string]string{"python": dockerfilePy, "golang": dockerfileGo}
	filePath := keys.FilePath + taskID + "/" + "Dockerfile"
	file, err := os.Create(filePath)
	
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(m[lang])
		
	if err != nil {
		panic(err)
	}
	file.Sync()
}