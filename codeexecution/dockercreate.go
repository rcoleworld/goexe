package codeexecution

import (
    "fmt"
    "os/exec"
)

func DockerPull(img string) {
	cmd := fmt.Sprintf("docker pull %s", img)
	pullCmd := exec.Command("bash", "-c", cmd)
	pullOut, err := pullCmd.Output()
	if err != nil {
		panic(err)
	} 
	fmt.Println(string(pullOut))
}

func DockerRun() string {
	buildCmd := exec.Command("bash", "-c", "docker build -t testapp ./codeexecution")
	// buildOut, err := 
	buildCmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(buildOut))
	runCmd := exec.Command("bash", "-c", "docker run testapp")
    runOut, err := runCmd.Output()
    if err != nil {
		panic(err)
    }
	fmt.Println(string(runOut))
	return string(runOut)
}