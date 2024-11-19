package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type InfoData struct {
	Id           primitive.ObjectID
	Header       string
	Description  string
	CoursesTitle string
	AboutCourses string
	Tag          string
}
