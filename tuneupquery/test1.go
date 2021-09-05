package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	arg1 := "SELECT * FROM EMPLOYEES"
	c := exec.Command("sudo", "./cmd_exec.sh", arg1)
	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}
	c.Stderr = stderr
	c.Stdout = stdout
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err, "|", stderr.String())
	} else {
		fmt.Println(stdout.String())
	}
	os.Exit(0)
}
