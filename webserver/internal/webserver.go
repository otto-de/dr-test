package webserver

import (
	"drtest/generated"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func generate(structName string, amount int) ([]interface{}, error) {
	strukt, err := generated.Generate(structName, amount)

	if err != nil {
		return nil, err
	}

	return strukt, nil
}

func getHandler(structName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		amount := 1 // TODO: use quantity from path var
		entities, err := generate(structName, amount)

		if err != nil {
			fmt.Println("ERROR generating:", err)
			os.Exit(1)
		}

		response, err := json.Marshal(entities)

		if err != nil {
			fmt.Println("MARSHAL ERROR", err)
		}

		_, writeErr := w.Write(response)

		if writeErr != nil {
			fmt.Println("WRITE ERROR", writeErr)
		}
	}
}

func populateRouter(router *mux.Router, structName string) {
	lowercased := strings.ToLower(structName)
	fmt.Println("Populating /" + lowercased)
	router.HandleFunc("/"+lowercased, getHandler(structName)).Methods("GET")
}

func getSchemaEntityName(schemaLocation string) string {
	jq := gojsonq.New().File(schemaLocation)
	res := jq.From("name").Get()
	return fmt.Sprint(res)
}

func StartServer(host string, port int, schemaLocations []string) {
	r := mux.NewRouter()

	for _, schemaLocation := range schemaLocations {
		populateRouter(r, getSchemaEntityName(schemaLocation))
	}

	addr := host + ":" + strconv.Itoa(port)

	fmt.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
