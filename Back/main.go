package main

import (
	"Back/API"
)

func main() {
	const portForWebsite = ":8080"
	API.Router.RunServer(portForWebsite)
}
