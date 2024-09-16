package transport

import (
	"context"
	"encoding/json" 
	"net/http"

	"github.com/gorilla/mux"
	kithttp "github.com/go-kit/kit/transport/http"

	"market-scoring/src/service"
)

// MakeHTTPHandler создает HTTP маршруты и эндпоинты
func MakeHTTPHandler(svc service.Service) http.Handler {
	r := mux.NewRouter()

	// Пример создания эндпоинта для обработки запроса
	getDataHandler := kithttp.NewServer(
		service.MakeGetDataEndpoint(svc),
		decodeGetDataRequest,
		encodeResponse,
	)

	r.Handle("/api/v1/data", getDataHandler).Methods("GET")
	return r
}

func decodeGetDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	sellerId := r.URL.Query().Get("sellerId")
	return service.GetDataRequest{SellerID: sellerId}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
