package service

import (
	"fmt"
	"log"
	"time"

	"market-scoring/src/config"
)


func IsDateOlderThanMonths(date time.Time, months int) bool {
	// Текущая дата
	currentDate := time.Now()

	// Дата, которая равна текущей дате минус количество месяцев
	thresholdDate := currentDate.AddDate(0, -months, 0)

	// Сравниваем переданную дату с пороговой датой
	return date.Before(thresholdDate)
}

func ScoringService(activeProducts int32, registrationDate time.Time, turnover float64, salesLastMonth int32) bool {

	// Инициализация конфигураций из файла
	err := config.InitConfigFromJSONFile("src/config/config.json")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Получение условий для скоринга для gRPC сервера из конфигурации
	scoringConditions := config.Config.ScoringConditions

	fmt.Println(activeProducts >= scoringConditions.ActiveProduct)
	fmt.Println(IsDateOlderThanMonths(registrationDate, 6))
	fmt.Println(turnover >= scoringConditions.Turnover)	
	fmt.Println(salesLastMonth >= scoringConditions.SalesLastMonth)


	return activeProducts >= scoringConditions.ActiveProduct && IsDateOlderThanMonths(registrationDate, 6) && turnover >= scoringConditions.Turnover && salesLastMonth >= scoringConditions.SalesLastMonth
}