package client

import (
	"context"
	"log"
	"net/http"
	"strconv"  // Для преобразования bool в строку
	"time"

	"google.golang.org/grpc"

	"market-scoring/src/config"
	"market-scoring/src/service"

	pb "market-scoring/src/proto"
)

func StartHTTPServer() {
	// Инициализация конфигураций из файла
	err := config.InitConfigFromJSONFile("src/config/config.json")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Проверка, что конфигурация загружена
	if config.Config == nil {
		log.Fatalf("Config is nil, initialization failed.")
	}

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		iin := r.URL.Query().Get("iin")
		bin := r.URL.Query().Get("bin")

		// Получаем адрес gRPC сервера из конфигурации
		grpcServerURL := config.Config.Urls.Grpc
		if grpcServerURL == "" {
			http.Error(w, "gRPC server URL is not configured", http.StatusInternalServerError)
			log.Println("gRPC server URL is missing in the config")
			return
		}

		// Устанавливаем соединение с gRPC-сервером
		conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure()) // Для тестов, используйте безопасное соединение в продуктиве
		if err != nil {
			http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
			log.Println("Failed to connect to gRPC server:", err)
			return
		}
		defer conn.Close()

		client := pb.NewDataServiceClient(conn)

		// Отправляем запрос gRPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := client.ProcessData(ctx, &pb.DataRequest{Iin: iin, Bin: bin})
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
	})

	// Получение URL для HTTP сервера из конфигурации
	serverURL := config.Config.Urls.Dev
	if serverURL == "" {
		log.Fatalf("HTTP server URL is not configured.")
	}

	// Запуск HTTP сервера на адресе, полученном из конфигурации
	log.Printf("HTTP server is running on %s\n", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, nil))
}
