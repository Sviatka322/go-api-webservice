package configs

import (
    "context"
    "fmt"
    "log"
    "time"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client  {
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        log.Fatal(err)
    }
  
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

var DB *mongo.Client = ConnectDB()

func GetPolicy(client *mongo.Client, url string) string {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    collection := client.Database("urlsdb").Collection("urls")
	urls, err := collection.Find(ctx, bson.M{"url": url})

	if err != nil {
		log.Fatal(err)
	}

	var db_url []bson.M

	if err = urls.All(context.TODO(), &db_url); err != nil {
		log.Fatal(err)
	}

	if len(db_url) > 0 {
		return fmt.Sprintf("%v", db_url[0]["policy"])
	} else {
		return ``
	}
}