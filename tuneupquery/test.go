package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	arg2 := "Test"
	cmd := exec.Command("bash", "-c", "sh ./cmd_exec.sh arg2")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
