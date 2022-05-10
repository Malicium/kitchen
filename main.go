package main

import (
	"os"
	"os/exec"
	"fmt"
	"bytes"
)


func main() {
	// router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.1.2"})

	// router.GET("/recipes", getRecipes)
	// router.Run("localhost:8080")
	RunCode("node", "hello.js")
}

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
		fmt.Println("An error occured: ", err)
	}

	process.Wait()
	fmt.Println(buf)
} 