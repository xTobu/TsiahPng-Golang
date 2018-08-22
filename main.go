package main

import (
	appstart "TsiahPng-Golang/AppStart"

	"google.golang.org/appengine"
)

func main() {
	appstart.RouteConfig()
	appengine.Main()
}
