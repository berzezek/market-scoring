package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	pb "market-scoring/src/proto"
)

func StartHTTPServer() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		iin := r.URL.Query().Get("iin")
		bin := r.URL.Query().Get("bin")

		// Устанавливаем соединение с gRPC-сервером
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
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
			return
		}

		// Форматируем вывод
		response := fmt.Sprintf(
			"Success: %s\nActive Products: %d\nRegistration Date: %s\nTurnover for 6 Months: %.2f\nSales Last Month: %d",
			res.Message,
			res.ActiveProducts,
			res.RegistrationDate.AsTime().Format(time.RFC3339),
			res.Turnover_6Months,
			res.SalesLastMonth,
		)

		// Ответ клиенту
		fmt.Fprintln(w, response)
	})

	fmt.Println("HTTP server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
