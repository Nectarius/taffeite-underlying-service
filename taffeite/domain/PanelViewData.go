package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type PanelViewData struct {
	Id             primitive.ObjectID
	Version        int
	Info           InfoData
	NavigationInfo NavigationInfoData
	Services       []TrainingDescriptionData
	Courses        []TrainingCourseData
}
