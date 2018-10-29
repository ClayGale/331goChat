package main

import (
	_ "database/sql"
	_ "html/template"
	_ "net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", mainHandler)

}
