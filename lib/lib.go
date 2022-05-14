package lib

import (
	"fmt"
	"bytes"
	"os"
	"os/exec"
)

func RunCode(cmd string, input string) {
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
		fmt.Println(err)
	}
	process.Wait()
	fmt.Println(buf)
} 