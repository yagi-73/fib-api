package main

import (
	"fibonacci-api/router"
)

func main() {
	e := router.NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}