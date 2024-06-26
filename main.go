package main

import (
	"validator/server"
)

func main() {
	server.StartServer(6969, "0.0.0.0")
}
