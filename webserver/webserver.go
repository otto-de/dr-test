package webserver

import "fmt"

func generateTestData(amount int, data interface{}) []interface{} {
	fmt.Printf("Got %d", data)
	return []interface{}{}
}
