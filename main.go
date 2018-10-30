package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3/blob/master/README.md
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
}

type User struct { //user struct for passing the username into the chat page
	username string
}

func welcome(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./login.gtpl") //sending login page
	t.Execute(w, nil)
}

var chatTemplate = template.Must(template.ParseFiles(os.getwd("./chat.tmpl")))

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// retrieving form data from inputs

	db, err := sql.Open("sqlite3", "./chat.db") //connecting to db
	checkErr(err)
	//creating users table if it doesnt exist already
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users ( name varchar(20), colour varchar(10))")
	checkErr(err)

	_, err = stmt.Exec() //running above sql
	checkErr(err)
	//inserting
	stmt, err = db.Prepare("INSERT INTO users(name, colour) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec(r.Form["name"], r.Form["colour"]) //running above sql with value parameters
	checkErr(err)
	fmt.Println(res)

	name := strings.Join(r.Form["name"], " ") //parsing the form input to a string
	data := &User{                            //passing the username to the chat page for future use
		username: name,
	}

	chatTemplate.Execute(w, data) //opening the chat page and passing the username for reference
	//t, _ := template.ParseFiles("chat.gtpl", data)
	//t.Execute(w, nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/login", login)
	http.HandleFunc("/sendMessage", sendMessage)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
