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
	SideName       string
	HostAddress    string
	DownloadRouter string
}

type Security struct {
	RetryCount  int
	Password    string
	AdminPasswd string
	AuthUA      bool
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	service := Service{
		FilePath: viper.GetString("service.FilePath"),
		ValidMin: viper.GetInt("service.ValidMin"),
	}

	customize := Customize{
		SideName:       viper.GetString("customize.SideName"),
		HostAddress:    viper.GetString("customize.HostAddress"),
		DownloadRouter: viper.GetString("customize.DownloadRouter"),
	}

	security := Security{
		RetryCount:  viper.GetInt("security.RetryCount"),
		Password:    viper.GetString("security.Password"),
		AdminPasswd: viper.GetString("security.AdminPassword"),
		AuthUA:      viper.GetBool("security.AuthUA"),
	}

	Cfg = Config{
		Port:      viper.GetInt("server.Port"),
		Service:   service,
		Customize: customize,
		Security:  security,
	}
}
