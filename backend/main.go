package main

import (
	"golang-notes/routes"
)

func main() {
	router := routes.Router()

	router.Run()
}
