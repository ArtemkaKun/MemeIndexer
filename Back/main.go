package main

import (
	"Back/API"
)

func main() {
	const PortForWebsite = ":8080"
	API.Router.RunServer(PortForWebsite)
}
