package main

import (
	"DBServer/API"
	"log"
)

func main() {
	log.Panic(API.Router.Run(":8001"))
}
