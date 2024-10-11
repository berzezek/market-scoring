package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configs структура для хранения всех конфигурационных данных
type Configs struct {
	Env       Env       `json:"env"`
	Urls      Urls      `json:"urls"`
	ScoringConditions ScoringConditions `json:"scoringConditions"`
}

// ScoringConditions структура для хранения критериев для всех пакетов
type ScoringConditions struct {
	XS PackageXS `json:"xs"`
	S  PackageS  `json:"s"`
	M  PackageM  `json:"m"`
	L  PackageL  `json:"l"`
	XL PackageXL `json:"xl"`
}

type Env struct {
	Mode string `json:"mode"`
}

type Urls struct {
	Dev      string `json:"dev"`
	Prod     string `json:"prod"`
	Grpc     string `json:"grpc"`
	GrpcProd string `json:"grpc_prod"`
}

type PackageXS struct {
	ActiveProductsAge int `json:"activeProductsAge"` // Товары активны и актуальны на сайте последние x дней
}

type PackageS struct {
	ActiveProductsAge                   int     `json:"activeProductsAge"`                   // Товары активны и актуальны на сайте последние x дней
	OrderCancellationBeforeConfirmation int     `json:"orderCancellationBeforeConfirmation"` // % отмен заказов до подтверждения наличия по вине продавца
	OrderCancellationDuringDelivery     int     `json:"orderCancellationDuringDelivery"`     // % отмен заказов на этапе доставки по вине продавца
	OrderConfirmationTimeViolation      int     `json:"orderConfirmationTimeViolation"`      // % заказов с нарушением временного регламента по подтверждению заказов
	OrderAssemblyTimeViolation          int     `json:"orderAssemblyTimeViolation"`          // % заказов с нарушением временного регламента по сборке заказов
	DeliveryTimeViolation               int     `json:"deliveryTimeViolation"`               // % нарушений сроков доставки (в случае доставки собственными силами)
	SellerRating                        float64 `json:"sellerRating"`                        // Рейтинг продавца на маркетплейсе
}

type PackageM struct {
	PackageS
}

type PackageL struct {
	PackageS
	Turnover          int  `json:"turnover"`          // Оборот продавца на маркетплейсе
	IsCreditBuyButton bool `json:"isCreditBuyButton"` // Наличие кнопки "Купить в кредит"
	IsQRCreditOffline bool `json:"isQRCreditOffline"` // Наличие QR-кода для оформления кредита оффлайн
}

type PackageXL struct {
	PackageL
}

var Config *Configs

// InitConfigFromJSONFile загружает конфигурацию из JSON файла
func InitConfigFromJSONFile(jsonFilePath string) error {
	var (
		filePath string
		configs  Configs
	)
	if os.Getenv("config") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filePath = pwd + "/" + jsonFilePath
	} else {
		filePath = os.Getenv("config")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configs)
	if err != nil {
		return err
	}

	Config = &configs
	return nil
}
