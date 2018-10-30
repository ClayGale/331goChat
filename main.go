package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3/blob/master/README.md
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
}

type User struct { //user struct for passing the username into the chat page
	name string
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// retrieving form data from inputs
		nameTemp := r.Form["name"]
		colourTemp := r.Form["colour"]
		name := strings.Join(nameTemp, " ")
		colour := strings.Join(colourTemp, " ")

		db, err := sql.Open("sqlite3", "./chat.db") //connecting to db
		//checkErr(err)

		stmt, err := db.Prepare("INSERT INTO users(name, colour) values(" + name + "," + colour + ")")
		//checkErr(err)

		res, err = stmt.Exec()
		//checkErr(err)

		data := &User{
			username: name,
		}

		t, _ := template.ParseFiles("chat.gtpl", data)
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/sendMessage", sendMessage)
}
