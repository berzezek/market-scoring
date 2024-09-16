package client

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"market-scoring/src/config"
	"market-scoring/src/service"

	pb "market-scoring/src/proto"
)

var grpcClient pb.DataServiceClient

func StartHTTPServer() {
	// Инициализация gRPC соединения один раз
	grpcServerURL := config.Config.Urls.Grpc // Доступ к конфигурации через config.Config
	if grpcServerURL == "" {
		log.Fatalf("gRPC server URL is not configured")
	}

	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	grpcClient = pb.NewDataServiceClient(conn)

	// Настройка HTTP маршрутов
	http.HandleFunc("/data", handleDataRequest)

	// Получение URL для HTTP сервера из конфигурации
	serverURL := config.Config.Urls.Dev // Доступ к URL из config.Config
	if serverURL == "" {
		log.Fatalf("HTTP server URL is not configured.")
	}

	// Запуск HTTP сервера
	log.Printf("HTTP server is running on %s\n", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, nil))
}

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	sellerId := r.URL.Query().Get("sellerId")

	// Отправляем запрос gRPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := grpcClient.ProcessData(ctx, &pb.DataRequest{SellerId: sellerId})
	if err != nil {
		http.Error(w, "Failed to process data", http.StatusInternalServerError)
		log.Println("Failed to process data:", err)
		return
	}

	// Преобразуем res.RegistrationDate из *timestamppb.Timestamp в time.Time
	registrationDate := res.RegistrationDate.AsTime()

	// Обработка ответа с помощью ScoringService
	scoringResult := service.ScoringService(res.ActiveProducts, registrationDate, res.Turnover, res.SalesLastMonth)

	// Преобразуем результат bool в строку
	response := strconv.FormatBool(scoringResult)

	// Ответ клиенту
	log.Println("Sending response to client")
	w.Write([]byte(response))
}
