package main

import (
		"os"
    "os/exec"
    "fmt"
    "bytes"
)

func main() {
	cmd := os.Getenv("CMD")
	input := os.Getenv("INPUT")
	process := exec.Command(cmd, input)
	stdin, err := process.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdin.Close()
	buf := new(bytes.Buffer)
	process.Stdout = buf
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	process.Wait()
	fmt.Println(buf)
}