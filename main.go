package main

import (
	_ "database/sql"
	_ "html/template"
	_ "net/http"
	_ "string"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3/blob/master/README.md
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/sendMessage", sendMessage)
}
