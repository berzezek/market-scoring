package main

import (
	"market-scoring/src/client"
	"market-scoring/src/server"
)

func main() {
	go server.StartGRPCServer() // Запуск gRPC-сервера в отдельной горутине
	client.StartHTTPServer()    // Запуск HTTP-сервера (REST API)
}
