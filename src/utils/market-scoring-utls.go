package utils

import "time"

func IsDateOlderThanMonths(date time.Time, months int) bool {
	// Текущая дата
	currentDate := time.Now()

	// Дата, которая равна текущей дате минус количество месяцев
	thresholdDate := currentDate.AddDate(0, -months, 0)

	// Сравниваем переданную дату с пороговой датой
	return date.Before(thresholdDate)
}