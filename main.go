package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	kullanici := "Kullanici: " + id

	message := API{kullanici}

	output, err := json.Marshal(message)

	checkError(err)

	fmt.Fprintf(w, string(output))
}
func main() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/hello/{id:[0-9]+}", Hello)

	http.Handle("/", gorillaRoute)

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error", err.Error())
		os.Exit(1)
	}
}
