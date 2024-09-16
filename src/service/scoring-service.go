package service

import (
	"time"

	"market-scoring/src/config"
	"market-scoring/src/utils"
)


func ScoringService(activeProducts int32, registrationDate time.Time, turnover float64, salesLastMonth int32) bool {
	scoringConditions := config.Config.ScoringConditions

	isActiveProductValid := activeProducts >= scoringConditions.ActiveProduct
	isRegistrationDateValid := utils.IsDateOlderThanMonths(registrationDate, 6)
	isTurnoverValid := turnover >= scoringConditions.Turnover
	isSalesLastMonthValid := salesLastMonth >= scoringConditions.SalesLastMonth

	return isActiveProductValid && isRegistrationDateValid && isTurnoverValid && isSalesLastMonthValid
}