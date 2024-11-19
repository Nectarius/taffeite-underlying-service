package services

import (
	entity "taffeite.com/taffeite-underlying-service/domain"
	"taffeite.com/taffeite-underlying-service/dto"
)

func mapDataToDto(source entity.TrainingDescriptionData) dto.TrainingDescriptionDto {
	return dto.TrainingDescriptionDto{
		Header:      source.Header,
		Description: source.Description,
		Image:       source.Image,
	}
}

func mapCourseDataToDto(source entity.TrainingCourseData) dto.TrainingCourseDto {
	return dto.TrainingCourseDto{
		Header:           source.Header,
		Description:      source.Description,
		AvailableOptions: source.AvailableOptions,
		Pricing:          source.Pricing,
	}
}

func ConvertToDto(data entity.PanelViewData) dto.PanelView {
	targetCollection := make([]dto.TrainingDescriptionDto, len(data.Services))
	for i, source := range data.Services {
		targetCollection[i] = mapDataToDto(source)
	}
	courseCollection := make([]dto.TrainingCourseDto, len(data.Courses))
	for i, source := range data.Courses {
		courseCollection[i] = mapCourseDataToDto(source)
	}
	return dto.PanelView{
		NavigationInfo: dto.NavigationInfoDto{
			Intro:          data.NavigationInfo.Intro,
			AboutTrainings: data.NavigationInfo.AboutTrainings,
			Pricing:        data.NavigationInfo.Pricing,
		},
		InfoDataSet: dto.InfoDataDto{
			Description:  data.Info.Description,
			Header:       data.Info.Header,
			AboutCourses: data.Info.AboutCourses,
			CoursesTitle: data.Info.CoursesTitle,
			Tag:          data.Info.Tag,
		},
		Services: targetCollection,
		Courses:  courseCollection,
	}
}
