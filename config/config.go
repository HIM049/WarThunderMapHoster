package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	NetWork
	Service
	Customize
	Security
	Secret
}

type NetWork struct {
	Port  int
	Https bool
}

type Service struct {
	MapDir      string
	ValidMin    int
	UpdateCheck bool
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

type Secret struct {
	SecretKey string
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	network := NetWork{
		Port:  viper.GetInt("server.Port"),
		Https: viper.GetBool("server.Https"),
	}

	service := Service{
		MapDir:      viper.GetString("service.MapDir"),
		ValidMin:    viper.GetInt("service.ValidMin"),
		UpdateCheck: viper.GetBool("service.UpdateCheck"),
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

	secret := Secret{
		SecretKey: viper.GetString("secret.SecretKey"),
	}

	Cfg = Config{
		NetWork:   network,
		Service:   service,
		Customize: customize,
		Security:  security,
		Secret:    secret,
	}
}
