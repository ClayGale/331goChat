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

	t := template.Must(template.ParseFiles("/go/src/github.com/ClayGale/331goChat/login.gtpl")) //sending login page
	t.Execute(w, nil)
}

//d, _ := os.Getwd()
var chatTemplate = template.Must(template.ParseFiles("/go/src/github.com/ClayGale/331goChat/chat.gohtml"))

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

func fetchPath() (d string) {
	d, _ = os.Getwd()
	return d
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

// this had to be hardcoded because docker is evil
const chatTemplateHTML = `<html>
<head>
  <link rel="stylesheet" type="text/css" href="chat.css">
  <link href='http://fonts.googleapis.com/css?family=Just+Another+Hand' rel='stylesheet' type='text/css'>
</head>
<body>
  <center><h2>Chat system using Golang</h2></center>
  <h3>Chatter Away...</h3>
  <center>
    <textarea rows="20" cols="100" style="resize: none; float: left; margin: 10px; text-align" id="dispMessages" name="dispMessages" disabled></textarea>
  </center>
  <center>
    <div class="main">
      <div id="container">
        <div class="chatBar">
          <form method="post">
            Message: <input type="text" class="message" placeholder="Type here to Chatter"><br><br>
            <input type="hidden" value="insertuserhere">
            <div class=button>
              <input type="submit" class="submit" value="Send Message">
            </div>
          </form>
        </div>
      </div>
    </div>
  </center>
</body>
</html>
    `
