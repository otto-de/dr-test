package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Webserver interface {
	GetRecordNames() []string
	GenerateEntity(recordName string, amount int) ([]interface{}, error)
}

func populateRouter(webserver Webserver, router *mux.Router, recordName string) {
	lowerCased := strings.ToLower(recordName)
	fmt.Println("Populating /" + lowerCased)
	router.HandleFunc("/"+lowerCased, func(w http.ResponseWriter, r *http.Request) {
		err := getHandler(webserver, recordName, 1, w)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}).Methods("GET")

	fmt.Println("Populating /" + lowerCased + "/{amount}")
	router.HandleFunc("/"+lowerCased+"/{amount:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		amount, err := amountFromRequest(r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = getHandler(webserver, recordName, amount, w)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}).Methods("GET")
}

func getHandler(webserver Webserver, recordName string, amount int, w http.ResponseWriter) error {
	handlerResponse, err := getResponse(webserver, recordName, amount)
	if err != nil {
		return err
	}

	response, err := json.Marshal(handlerResponse.Body)
	if err != nil {
		return err
	}

	// Headers must be written before calling w.Write,
	// since it will call WriteHeader if it hasn't been
	// called before
	for name, value := range handlerResponse.Headers {
		w.Header().Set(name, value)
	}

	_, err = w.Write(response)
	if err != nil {
		return err
	}

	return nil
}

func amountFromRequest(r *http.Request) (int, error) {
	params := mux.Vars(r)
	amount := 1
	if amountStr, found := params["amount"]; found {
		parsedAmount, err := strconv.Atoi(amountStr)
		if err != nil {
			return 0, err
		}
		amount = parsedAmount
	}
	return amount, nil
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
