package main

import (
    "log"
    "api-webservice/configs"
	"api-webservice/routes"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	configs.ConnectDB()

	routes.UserRoute(router)
	log.Fatal(http.ListenAndServe(":10000", router))
}