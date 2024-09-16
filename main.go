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
	Mode := config.Config.Env.Mode
	
	if (Mode=="dev") {
		fmt.Println("Mode is development, Start server...")
		go server.StartGRPCServer()
	}
	client.StartHTTPServer()
}
