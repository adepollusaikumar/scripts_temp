package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	Query  string
	Source string
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	q := req.FormValue("query")
	c := exec.Command("sudo", "./cmd_exec.sh", q)
	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}
	c.Stderr = stderr
	c.Stdout = stdout
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err, "|", stderr.String())
	} else {
		fmt.Println(stdout.String())
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{stdout.String(), q})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
