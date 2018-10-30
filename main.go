package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3/blob/master/README.md
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
}

type User struct { //user struct for passing the username into the chat page
	username string
}

var chatTemplate = template.Must(template.ParseFiles("chat.gtpl"))

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// retrieving form data from inputs
		//nameTemp := r.Form["name"]
		//colourTemp := r.Form["colour"]
		//name := strings.Join(nameTemp, " ")
		//colour := strings.Join(colourTemp, " ")

		db, err := sql.Open("sqlite3", "./chat.db") //connecting to db
		checkErr(err)
		//creating users table if it doesnt exist already
		stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users ( name varchar(20), colour varchar(10))")
		checkErr(err)

		_, err := stmt.Exec() //running above sql
		checkErr(err)
		//inserting
		stmt, err := db.Prepare("INSERT INTO users(name, colour) values(?,?)")
		checkErr(err)

		res, err := stmt.Exec(r.Form["name"], r.Form["colour"]) //running above sql with value parameters
		checkErr(err)
		fmt.Println(res)

		data := &User{
			username: name,
		}

		chatTemplate.Execute(w, data) //opening the chat page and passing the username for reference
		//t, _ := template.ParseFiles("chat.gtpl", data)
		//t.Execute(w, nil)
	}
}

func checkErr(string err) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/sendMessage", sendMessage)
}
