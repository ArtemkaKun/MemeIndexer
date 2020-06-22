package main

import (
	"DBServer/API"
	"log"
)

func main() {
	const PortForServer string = ":8001"
	log.Panic(API.Router.Run(PortForServer))
}
