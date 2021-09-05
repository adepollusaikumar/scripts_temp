package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"

		name := req.FormValue(key)
		c := exec.Command("sudo", "./cmd_exec.sh", name)
		stderr := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		c.Stderr = stderr
		c.Stdout = stdout
		if err := c.Run(); err != nil {
			fmt.Println("Error: ", err, "|", stderr.String())
		} else {
			/*	fmt.Println(stdout.String())*/
			fmt.Fprintf(w, "%s\n", stdout.String())
		}

		val := req.
			fmt.Println("value: ", val)
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST">
		 <input type="text" name="q">
		 <input type="submit">
		</form>`+val)
	})
	http.ListenAndServe(":8080", nil)
}
