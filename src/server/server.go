package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	"market-scoring/src/config"
	pb "market-scoring/src/proto"
)

type server struct {
	pb.UnimplementedDataServiceServer
}

// Генерация случайного числа в заданном диапазоне [min, max]
func randomIntInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Генерация случайного числа с плавающей точкой в заданном диапазоне [min, max]
func randomFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// ProcessData обрабатывает входящий запрос и возвращает динамический ответ в зависимости от входных данных.
func (s *server) ProcessData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	// Устанавливаем seed для генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Генерируем случайные значения в указанных диапазонах
	activeProductsAge := int32(randomIntInRange(90, 200))                                // Число от 90 до 200
	orderCancellationBeforeConfirmation := int32(randomIntInRange(5, 15))                // От 5 до 15
	orderCancellationDuringDelivery := int32(randomIntInRange(5, 15))                    // От 5 до 15
	orderConfirmationTimeViolation := int32(randomIntInRange(5, 15))                     // От 5 до 15
	orderAssemblyTimeViolation := int32(randomIntInRange(3, 10))                         // От 3 до 10
	deliveryTimeViolation := int32(randomIntInRange(3, 10))                              // От 3 до 10
	sellerRating := randomFloatInRange(4.0, 5.0)                                         // Рейтинг от 4.0 до 5.0
	turnover := int32(randomIntInRange(3_000_000, 10_000_000))                           // Оборот от 3 млн до 10 млн
	isCreditBuyButton := rand.Intn(2) == 0                                               // Случайное булево значение
	isQRCreditOffline := rand.Intn(2) == 0                                               // Случайное булево значение

	// Логируем запрос и генерируем ответ
	fmt.Printf("Received Seller ID: %s\n", req.SellerId)

	// Возвращаем динамически сгенерированный ответ
	return &pb.DataResponse{
		ActiveProductsAge:                   activeProductsAge,
		OrderCancellationBeforeConfirmation: orderCancellationBeforeConfirmation,
		OrderCancellationDuringDelivery:     orderCancellationDuringDelivery,
		OrderConfirmationTimeViolation:      orderConfirmationTimeViolation,
		OrderAssemblyTimeViolation:          orderAssemblyTimeViolation,
		DeliveryTimeViolation:               deliveryTimeViolation,
		SellerRating:                        sellerRating,
		Turnover:                            turnover,
		IsCreditBuyButton:                   isCreditBuyButton,
		IsQRCreditOffline:                   isQRCreditOffline,
	}, nil
}

func StartGRPCServer() {
	grpcServerAddress := config.Config.Urls.Grpc

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
