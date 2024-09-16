# Market scoring service

## Instalation


### Protobuf
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH=$PATH:$(go env GOPATH)/bin

protoc --go_out=. --go-grpc_out=. src/proto/data.proto
```

### Scoring options

#### Для установки условий скорнинга см. config/config.json -> scoringConditions где
* activeProduct - Минимальное кол-во активных товаров
* RegistrationDate - Минимальное кол-во месяцев прошедших с начала регистрации
* Turnover - Минимальный оборот за 6 месяцев
* SalesLastMonth - Кол-во продаж за последний месяц


### Условия по умолчанию
1. Длительность работы на маркетплейсе: Кредит предоставляется только мерчантам, которые работают на платформе не менее 6 месяцев.

2. Количество активных товаров: Мерчант должен иметь в наличии не менее 10 активных товаров, которые регулярно обновляются и пополняются

3. Количество проданных товаров: Мерчант должен продать минимум 5 товаров за последний месяц

4. Кол-во продаж: тут за последний 6 мес мерчант должен продать на 6 млн выше


### Request
```sh
curl "http://localhost:8080/data?iin=123&bin=456"
```

