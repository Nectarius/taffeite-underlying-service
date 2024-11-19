package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"log"
	"taffeite.com/taffeite-underlying-service/conf"
	entity "taffeite.com/taffeite-underlying-service/domain"
)

type InfoDataRepository struct {
	Conf conf.MongoConf
}

func NewInfoDataRepository(Conf conf.MongoConf) *InfoDataRepository {
	return &InfoDataRepository{Conf: Conf}
}

func (r *InfoDataRepository) GetInfoData() entity.InfoData {
	var client = r.Conf.MongoClient
	collection := client.Database("taffeite").Collection("info-data")
	result := collection.FindOne(context.Background(), bson.M{})

	var infoData = entity.InfoData{}
	var err = result.Decode(&infoData)
	if err != nil {
		log.Fatal(err)
	}

	return infoData
}
