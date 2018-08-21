package main

import (
	appstart "TsiahPng-Golang/AppStart"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"
)

var db *sql.DB

func main() {

	appstart.RouteConfig()
	appengine.Main()
}
