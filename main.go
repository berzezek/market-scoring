package main

import (
	"fmt"
	"log"

	"market-scoring/src/client"
	"market-scoring/src/config"
	"market-scoring/src/server"
)

func main() {
	// Инициализация конфигураций из файла
	err := config.InitConfigFromJSONFile("src/config/config.json")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Получение mode из конфигурации
	mode := config.Config.Env.Mode
	var grpcServerURL string

	if mode == "dev" {
		fmt.Println("Mode is development, starting gRPC server...")
		go server.StartGRPCServer() // Запуск gRPC сервера в фоновом режиме
		grpcServerURL = config.Config.Urls.Grpc // Используем URL для разработки
	} else {
		fmt.Println("Mode is production, starting HTTP server...")
		grpcServerURL = config.Config.Urls.GrpcProd // Используем URL для продакшена
	}

	// Запуск HTTP сервера с gRPC URL и передачей сервиса
	client.StartHTTPServer(grpcServerURL)
}
