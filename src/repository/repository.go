package repository

import (
	"context"
	"google.golang.org/grpc"
	"market-scoring/src/proto"
)

// Repository интерфейс для работы с внешними сервисами (gRPC)
type Repository interface {
	GetExternalData(ctx context.Context, sellerId string) (*proto.DataResponse, error)
}

// grpcRepository структура для работы с gRPC
type grpcRepository struct {
	client proto.DataServiceClient
}

// NewRepository создает новый репозиторий для работы с gRPC
func NewRepository(conn *grpc.ClientConn) Repository {
	client := proto.NewDataServiceClient(conn)
	return &grpcRepository{client: client}
}

// GetExternalData реализует метод запроса данных через gRPC
func (r *grpcRepository) GetExternalData(ctx context.Context, sellerId string) (*proto.DataResponse, error) {
	req := &proto.DataRequest{SellerId: sellerId}
	return r.client.ProcessData(ctx, req)
}
