package webserver

import (
	"drtest/generated"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getHandler(recordName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		amount := 1 // TODO: use quantity from path var
		entities, err := generated.Generate(recordName, amount)

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

func populateRouter(router *mux.Router, recordName string) {
	lowercased := strings.ToLower(recordName)
	fmt.Println("Populating /" + lowercased)
	router.HandleFunc("/"+lowercased, getHandler(recordName)).Methods("GET")
}

func StartServer(host string, port int) {
	r := mux.NewRouter()

	for _, recordName := range generated.GetRecordNames() {
		populateRouter(r, recordName)
	}

	addr := host + ":" + strconv.Itoa(port)

	fmt.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
