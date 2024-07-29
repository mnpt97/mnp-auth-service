package main

import (
	"fmt"
	authservice "mnp-auth-service/auth_service"
	dbservice "mnp-auth-service/db_service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	simpleDb := dbservice.NewSimpleDb("/Users/mnp/workspace/mnp-auth-service-go/data/min-user.json")
	router := mux.NewRouter()

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { authservice.LoginHandler(w, r, simpleDb) })
	router.HandleFunc("/secret", authservice.Grant(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Granted")
	}))
	fmt.Println("Starting the server")
	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}

}
