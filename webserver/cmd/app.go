package main

import webserver "drtest/webserver/internal"

func main() {
	webserver.StartServer("localhost", 8080)
}