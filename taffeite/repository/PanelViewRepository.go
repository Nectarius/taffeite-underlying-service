package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"taffeite.com/taffeite-underlying-service/conf"
	entity "taffeite.com/taffeite-underlying-service/domain"
)

type PanelViewRepository struct {
	Conf conf.MongoConf
}

func NewPanelViewRepository(Conf conf.MongoConf) *PanelViewRepository {
	return &PanelViewRepository{Conf: Conf}
}

func (r *PanelViewRepository) GetPanelViewData() entity.PanelViewData {
	var client = r.Conf.MongoClient
	collection := client.Database("taffeite").Collection("panel-view-data")
	var filter = bson.M{"version": 3}
	result := collection.FindOne(context.Background(), filter)

	var panelView = entity.PanelViewData{}
	var err = result.Decode(&panelView)
	if err != nil {
		log.Fatal(err)
	}

	return panelView
}

func (r *PanelViewRepository) InsertPanelViewData(panelView entity.PanelViewData) primitive.ObjectID {
	var client = r.Conf.MongoClient
	collection := client.Database("taffeite").Collection("panel-view-data")
	one, err := collection.InsertOne(context.Background(), panelView)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return one.InsertedID.(primitive.ObjectID)
}
