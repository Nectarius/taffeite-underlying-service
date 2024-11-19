package services

import (
	"taffeite.com/taffeite-underlying-service/dto"
)

func GetDefaultInfoData() dto.InfoDataDto {
	var infoData = dto.InfoDataDto{}
	infoData.Header = "Здравствуйте меня зовут Ольга Молодцова"
	infoData.Description = "Хореограф/Тренер и балерина. Преподаватель балета и латиноамериканских танцев более 12 лет. Автор и преподаватель курса 'Постройней за 30 дней'. Более 1500 довольных клиентов обучила растяжке, красивой осанке и женственности."
	infoData.AboutCourses = "1. Убрать лишний вес\n2. Раскрыть женскую сексуальность\n3. Научиться танцевать\n4. Победить депрессию и апатию\n5. Эффективные тренировки без спортзала\n6. Понять своё предназначение и "
	infoData.CoursesTitle = "Зачем Вы здесь ?"
	return infoData
}
