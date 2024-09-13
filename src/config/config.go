package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	Env  Env  `json:"env"`
	Urls Urls `json:"urls"`
}

type Env struct {
	Mode string `json:"mode"`
}

type Urls struct {
	Dev  string `json:"dev"`
	Prod string `json:"prod"`
	Grpc string `json:"grpc"`
}

var Config *Configs

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
