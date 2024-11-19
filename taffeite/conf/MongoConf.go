package conf

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"os"
)

type MongoConf struct {
	MongoUri    string
	MongoClient *mongo.Client
}

func NewMongoConf() *MongoConf {
	data, err := os.ReadFile("resources/conf.yaml")
	if err != nil {
		panic(err)
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	mongoUrlPath, ok := config["mongo-url-path"].(string)
	if !ok {
		fmt.Println("Property not found")
	} else {
		fmt.Println("Value:", mongoUrlPath)
	}

	clientOptions := options.Client().ApplyURI(mongoUrlPath)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConf{
		MongoUri:    mongoUrlPath,
		MongoClient: client,
	}
}

func (r *MongoConf) Clear() {
	err := r.MongoClient.Disconnect(context.Background())
	log.Println("MongoConf Disconnected ")
	if err != nil {
		log.Fatal(err)
	}

}
