package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	"market-scoring/src/config"
	"market-scoring/src/proto"
	"market-scoring/src/utils"
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

	// Логируем gRPC ответ
	fmt.Printf("Received gRPC response: %+v\n", grpcRes)

	// Вызываем ScoringService для определения числового значения категории
	packageLevel := ScoringService(grpcRes)

	// Возвращаем числовое значение категории в ответе
	return GetDataResponse{
		Message: packageLevel,
	}, nil
}

// ScoringService анализирует ответ gRPC и возвращает числовое значение категории (от 0 до 5)
func ScoringService(grpcRes *proto.DataResponse) int {
	// Загружаем условия скоринга из конфигурации
	conditions := config.Config.ScoringConditions

	// Сравниваем gRPC ответ с условиями для каждого пакета
	switch {
	case utils.MatchesPackageXL(grpcRes, conditions.XL):
		return 5 // XL
	case utils.MatchesPackageL(grpcRes, conditions.L):
		return 4 // L
	case utils.MatchesPackageM(grpcRes, conditions.M):
		return 3 // M
	case utils.MatchesPackageS(grpcRes, conditions.S):
		return 2 // S
	case utils.MatchesPackageXS(grpcRes, conditions.XS):
		return 1 // XS
	default:
		return 0 // Если ни один пакет не подошел
	}
}

// GetDataRequest структура запроса
type GetDataRequest struct {
	SellerID string
}

// GetDataResponse структура ответа
type GetDataResponse struct {
	Message int `json:"message"`
}

// MakeGetDataEndpoint создает эндпоинт для метода GetData
func MakeGetDataEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDataRequest)
		return svc.GetData(ctx, req)
	}
}
