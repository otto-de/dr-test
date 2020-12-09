package main

import (
	webserver "drtest/webserver/internal"
	"log"
	"os"
)

func main() {
	log.Print("Used schema locations:", os.Args[1:])
	webserver.StartServer("localhost", 8080, os.Args[1:])
}
