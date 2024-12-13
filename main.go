package main

import (
	"GroupieTracker/routes"
	"GroupieTracker/templates"
)

func main() {
	templates.InitTmpl()
	routes.InitServ()
}
