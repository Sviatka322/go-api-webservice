package routes

import (
	 "github.com/gorilla/mux"
	 "api-webservice/controllers"
	 "github.com/swaggo/http-swagger"
	 _"api-webservice/docs"
)

func UserRoute(router *mux.Router)  {
    router.HandleFunc("/urlinfo/1/{url}/{search_query}", controllers.GetPolicy()).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}