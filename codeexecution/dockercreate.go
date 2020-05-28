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

func DockerRun(dir string) string {
	buildCmd := exec.Command("bash", "-c", fmt.Sprintf("docker build -t %s ./codeexecution/%s", dir, dir))
	fmt.Println(buildCmd)
	buildOut, _ := buildCmd.Output()
	fmt.Println(string(buildOut))
	runCmd := exec.Command("bash", "-c", fmt.Sprintf("docker run %s", dir))
	fmt.Println(runCmd)
    runOut, err := runCmd.Output()
    if err != nil {
		fmt.Println(err)
    }
	fmt.Println(string(runOut))
	return string(runOut)
}