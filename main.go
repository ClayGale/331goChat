package main

import (
	_ "database/sql"
	_ "fmt"
	_ "html/template"
	_ "net/http"
	_ "strings"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3/blob/master/README.md
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.form)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// send inputs to chat page TODO
	}
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/sendMessage", sendMessage)
}
