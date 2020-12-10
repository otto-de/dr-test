package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Webserver interface {
	GetRecordNames() []string
	GenerateEntity(recordName string, amount int) (interface{}, error)
}

func populateRouter(webserver Webserver, router *mux.Router, recordName string) {
	lowerCased := strings.ToLower(recordName)
	fmt.Println("Populating /" + lowerCased)
	router.HandleFunc("/"+lowerCased, func(w http.ResponseWriter, r *http.Request) {
		amount := 1 // TODO: use quantity from path var
		handlerResponse, err := getResponse(webserver, recordName, amount)

		if err != nil {
			fmt.Println("ERROR generating:", err)
			os.Exit(1)
		}

		response, err := json.Marshal(handlerResponse.Body)

		if err != nil {
			fmt.Println("MARSHAL ERROR", err)
		}

		_, writeErr := w.Write(response)

		if writeErr != nil {
			fmt.Println("WRITE ERROR", writeErr)
		}

		for name, value := range handlerResponse.Headers {
			w.Header().Set(name, value)
		}
	}).Methods("GET")
}

func Start(server Webserver, host string, port int) {
	r := mux.NewRouter()

	for _, recordName := range server.GetRecordNames() {
		populateRouter(server, r, recordName)
	}

	addr := host + ":" + strconv.Itoa(port)

	fmt.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
