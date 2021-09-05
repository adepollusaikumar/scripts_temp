package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	/*fmt.Fprintf(w, "POST request successful")*/
	name := r.FormValue("name")
	c := exec.Command("sudo", "./cmd_exec.sh", name)
	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}
	c.Stderr = stderr
	c.Stdout = stdout
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err, "|", stderr.String())
	} else {
		fmt.Println(stdout.String())
	}

	n, err := fmt.Fprintf(w, "%s\n", stdout.String())
	fmt.Println(err)
	g := strconv.Itoa(n)
	/*fmt.Fprintf(w, "%s\n", stdout.String())*/
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 80\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
