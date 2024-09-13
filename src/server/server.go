package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	
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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataServiceServer(grpcServer, &server{})
	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
