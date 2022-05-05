package main

import (
	"fmt"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

type Url struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Policy string `json:"policy"`
}

func main() {
	url_1 := Url{Id: 1, Url: "google.com", Policy: "blocked"}
	url_2 := Url{Id: 2, Url: "facebook.com", Policy: "blocked"}
	url_3 := Url{Id: 3, Url: "cisco.com", Policy: "unblocked"}
	url_4 := Url{Id: 4, Url: "netflix.com", Policy: "unblocked"}
	var urls_data []Url
	urls_data = append(urls_data, url_1)
	urls_data = append(urls_data, url_2)
	urls_data = append(urls_data, url_3)
	urls_data = append(urls_data, url_4)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
        log.Fatal(err)
    }
	urls := client.Database("urlsdb").Collection("urls")

	for _, element := range urls_data {
		urlsResult, err := urls.InsertOne(context.TODO(), element)

		fmt.Println(element)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(urlsResult)
	}

}