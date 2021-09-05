package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("q")
	c := exec.Command("sudo", "./cmd_exec.sh", name)
	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}
	c.Stderr = stderr
	c.Stdout = stdout
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err, "|", stderr.String())
	} else {
		fmt.Println(stdout.String())
		fmt.Fprintf(w, "%s\n", stdout.String())
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
    <textarea type="text" name="q" id="KUNDE" placeholder="Paste select Q" value="" rows="25"></textarea>
    <input type="Submit" value="Go>>" >
	<textarea type="text" name="r" id="KUNDE"  value="" rows="25"></textarea>
    </form>
    <hr>`+stdout.String())
}
