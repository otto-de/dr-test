package main

import (
	"drtest/webserver/internal"
	webserver "drtest/webserver/pkg"
)

func main() {
	webserver.Start(internal.ProductionWebserver{}, "127.0.0.1", 8080)
}
