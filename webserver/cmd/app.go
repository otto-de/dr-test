package main

import (
	webserver "drtest/webserver/internal"
)

func main() {
	webserver.StartServer("127.0.0.1", 8080)
}
