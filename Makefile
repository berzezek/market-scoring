# Переменные
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

# Установка зависимостей
install:
	@echo "Устанавливаем необходимые зависимости..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Экспорт пути
export_path:
	@echo "Добавляем $(go env GOPATH)/bin в PATH..."
	export PATH=$$PATH:$(go env GOPATH)/bin

# Генерация protobuf файлов
proto:
	@echo "Генерация protobuf файлов..."
	protoc --go_out=. --go-grpc_out=. src/proto/data.proto

# Компиляция и запуск проекта
run:
	@echo "Запуск проекта..."
	go run main.go

# Команда для полной сборки и запуска
all: install export-path proto run

# Очистка сгенерированных файлов (если необходимо)
clean:
	@echo "Удаление сгенерированных protobuf файлов..."
	rm -f src/proto/*.pb.go
