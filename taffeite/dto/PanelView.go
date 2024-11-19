package dto

type PanelView struct {
	NavigationInfo NavigationInfoDto
	InfoDataSet    InfoDataDto
	Services       []TrainingDescriptionDto
	Courses        []TrainingCourseDto
}
