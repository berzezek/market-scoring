package client

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"market-scoring/src/config"
	"market-scoring/src/transport"
	"market-scoring/src/service"
	pb "market-scoring/src/proto"
)

func StartHTTPServer(grpcServerURL string) {
	if grpcServerURL == "" {
		log.Fatalf("gRPC server URL is not configured")
	}

	// Инициализация gRPC соединения
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewDataServiceClient(conn)

	// Инициализация сервиса с gRPC клиентом
	svc := service.NewService(grpcClient)

	// Создание HTTP маршрутов с использованием transport.MakeHTTPHandler
	handler := transport.MakeHTTPHandler(svc)

	// Получение URL для HTTP сервера из конфигурации
	serverURL := config.Config.Urls.Dev
	if serverURL == "" {
		log.Fatalf("HTTP server URL is not configured.")
	}

	// Запуск HTTP сервера
	log.Printf("HTTP server is running on %s\n", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, handler))
}
