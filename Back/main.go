package main

import (
	"Back/API"
	"log"
)

func main() {
	log.Panic(API.Router.Run(":8080"))
}
