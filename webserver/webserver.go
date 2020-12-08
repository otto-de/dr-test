package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func getHandler(structName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type DemoStruct struct {
			Name string
		}

		fmt.Println("STRUCTNAME", structName)

		response, err := json.Marshal(DemoStruct{structName})

		if err != nil {
			fmt.Println("MARSHAL ERROR", err)
		}

		_, writeErr := w.Write(response)

		if writeErr != nil {
			fmt.Println("WRITE ERROR", writeErr)
		}
	}
}

func populateRouter(router *mux.Router, structNames []string) {
	for _, structName := range structNames {
		router.HandleFunc("/" + structName, getHandler(structName)).Methods("GET")
	}
}

func StartServer(host string, port int) {
	r := mux.NewRouter()
	populateRouter(r, []string{"person"})
	addr := host + ":" + strconv.Itoa(port)
	fmt.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
