package main

import (
	"gintest/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":9999")
}
