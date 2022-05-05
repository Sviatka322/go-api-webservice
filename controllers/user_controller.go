package controllers

import (
    "encoding/json"
    "api-webservice/configs"
    "net/http"
	"regexp"
	"github.com/gorilla/mux"
	"fmt"
)

// GetPolicyByURL
// @Summary Get URL policy
// @Tags policy
// @Accept json
// @Produce json
// @Param url path string true "url"
// @Param search_query path string true "query"
// @Success 200 {object} string
// @Success 204 {object} string
// @Failure 404 {object} string
// @Router /urlinfo/1/{url}/{search_query} [get]
func GetPolicy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key_url := vars["url"]
		fmt.Println(key_url)
		re := regexp.MustCompile(`^[a-z][a-z0-9+\-.]*`)
	
		matchall := re.FindAllString(key_url,-1)

		policy := configs.GetPolicy(configs.DB, matchall[0])
	
		if policy != `` {
			data, _ := json.Marshal(policy)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(data))

		} else {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("☄ HTTP status code returned! 204 No Content"))
		}
	}
}