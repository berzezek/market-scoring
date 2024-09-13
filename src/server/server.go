package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"market-scoring/src/config"  // Импортируем конфигурации
	pb "market-scoring/src/proto"
)

type server struct {
	pb.UnimplementedDataServiceServer
}

func (s *server) ProcessData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	// Примерные данные, которые сервер возвращает
	activeProducts := 150
	registrationDate := timestamppb.New(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
	turnover := float32(123456.78)
	salesLastMonth := 320

	// Логируем запрос
	fmt.Printf("Received IIN: %s, BIN: %s\n", req.Iin, req.Bin)

	// Возвращаем ответ с данными
	return &pb.DataResponse{
		Message:          "Data processed",
		Success:          true,
		ActiveProducts:   int32(activeProducts),
		RegistrationDate: registrationDate,
		Turnover_6Months:  turnover,
		SalesLastMonth:   int32(salesLastMonth),
	}, nil
}

func StartGRPCServer() {
	// Инициализация конфигураций из файла
	err := config.InitConfigFromJSONFile("src/config/config.json")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Получение URL для gRPC сервера из конфигурации
	grpcServerAddress := config.Config.Urls.Grpc

	// Запуск gRPC сервера на адресе, полученном из конфигурации
	lis, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataServiceServer(grpcServer, &server{})
	fmt.Printf("gRPC server is running on %s\n", grpcServerAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
