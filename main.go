package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	Firstname string "json:firstname"
	Lastname  string "json:lastname"
	Age       int    "json:age"
}

func main() {
	apiRoot := "/api"
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{
				ID:        0,
				Firstname: "zafer",
				Lastname:  "caliskan",
				Age:       24,
			},
			User{
				ID:        1,
				Firstname: "Gizem",
				Lastname:  "Hamamcioglu",
				Age:       22,
			},
		}

		output, err := json.Marshal(users)

		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error", err.Error())
		os.Exit(1)
	}
}
