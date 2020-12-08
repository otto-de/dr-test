package main

import webserver "drtest/internal"

func main() {
	webserver.StartServer("localhost", 8080)
}