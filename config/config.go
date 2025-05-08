package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Port int
	Service
	Customize
	Security
}

type Service struct {
	FilePath string
	ValidMin int
}

type Customize struct {
	SideName string
}

type Security struct {
	RetryCount  int
	Password    string
	AdminPasswd string
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	service := Service{
		FilePath: viper.GetString("service.filepath"),
		ValidMin: viper.GetInt("service.validmin"),
	}

	customize := Customize{
		SideName: viper.GetString("customize.sidename"),
	}

	security := Security{
		RetryCount:  viper.GetInt("security.retrycount"),
		Password:    viper.GetString("service.password"),
		AdminPasswd: viper.GetString("service.adminpassword"),
	}

	Cfg = Config{
		Port:      viper.GetInt("server.port"),
		Service:   service,
		Customize: customize,
		Security:  security,
	}
}
