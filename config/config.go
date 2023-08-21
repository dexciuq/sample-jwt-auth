package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host string
	Port int
	DB   struct {
		Postgres postgres
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
}

type postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadConfiguration() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file: %v\n", err)
		return nil, err
	}

	return &Config{
		Host: viper.GetString("server.host"),
		Port: viper.GetInt("server.port"),
		DB: struct{ Postgres postgres }{
			postgres{
				Host:     viper.GetString("db.postgres.host"),
				Port:     viper.GetInt("db.postgres.port"),
				User:     viper.GetString("db.postgres.user"),
				Password: viper.GetString("db.postgres.password"),
				DBName:   viper.GetString("db.postgres.dbname"),
			},
		},
		SMTP: struct {
			Host     string
			Port     int
			Username string
			Password string
			Sender   string
		}{
			Host:     viper.GetString("smtp.host"),
			Port:     viper.GetInt("smtp.port"),
			Username: viper.GetString("smtp.username"),
			Password: viper.GetString("smtp.password"),
			Sender:   viper.GetString("smtp.sender"),
		},
	}, nil
}
