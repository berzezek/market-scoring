package service

import (
	"context"
	"time"

	"market-scoring/src/config"
	"market-scoring/src/proto"
	"market-scoring/src/utils"

	"github.com/go-kit/kit/endpoint"
)

// Service описывает интерфейс нашего сервиса
type Service interface {
	GetData(ctx context.Context, req GetDataRequest) (GetDataResponse, error)
}

// service структура, реализующая интерфейс Service
type service struct {
	grpcClient proto.DataServiceClient // Добавим grpcClient в структуру service
}

// NewService создает новый сервис с gRPC клиентом
func NewService(grpcClient proto.DataServiceClient) Service {
	return &service{
		grpcClient: grpcClient, // Инициализация grpc клиента
	}
}

// GetData реализует бизнес-логику для получения данных через gRPC
func (s *service) GetData(ctx context.Context, req GetDataRequest) (GetDataResponse, error) {
	// Отправляем gRPC запрос для получения данных
	grpcRes, err := s.grpcClient.ProcessData(ctx, &proto.DataRequest{SellerId: req.SellerID})
	if err != nil {
		return GetDataResponse{}, err
	}

	// Преобразуем gRPC ответ в формат Go
	registrationDate := grpcRes.RegistrationDate.AsTime()

	// Обработка ответа с помощью ScoringService
	scoringResult := ScoringService(grpcRes.ActiveProducts, registrationDate, grpcRes.Turnover, grpcRes.SalesLastMonth)

	// Возвращаем результат как булево значение
	return GetDataResponse{
		Message: scoringResult, // Результат скоринга в виде булевого значения
	}, nil
}


// ScoringService выполняет анализ данных
func ScoringService(activeProducts int32, registrationDate time.Time, turnover float64, salesLastMonth int32) bool {
	scoringConditions := config.Config.ScoringConditions

	isActiveProductValid := activeProducts >= scoringConditions.ActiveProduct
	isRegistrationDateValid := utils.IsDateOlderThanMonths(registrationDate, 6)
	isTurnoverValid := turnover >= scoringConditions.Turnover
	isSalesLastMonthValid := salesLastMonth >= scoringConditions.SalesLastMonth

	return isActiveProductValid && isRegistrationDateValid && isTurnoverValid && isSalesLastMonthValid
}

// GetDataRequest структура запроса
type GetDataRequest struct {
	SellerID string
}

// GetDataResponse структура ответа
type GetDataResponse struct {
	Message bool `json:"message"`
}

// MakeGetDataEndpoint создает эндпоинт для метода GetData
func MakeGetDataEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDataRequest)
		return svc.GetData(ctx, req)
	}
}
