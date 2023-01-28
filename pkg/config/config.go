package config

import (
	"github.com/spf13/viper"
)

type DB struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name string `mapstructure:"NAME"`
}

type RabbitMQ struct {
	Host string `mapstructure:"R_HOST"`
	Port string `mapstructure:"R_PORT"`
	User string `mapstructure:"R_USER"`
	Password string `mapstructure:"R_PASSWORD"`
}

type Config struct {
	Port string `mapstructure:"PORT"`
	DB DB `mapstructure:"DB"`
	RabbitMQ RabbitMQ `mapstructure:"RABBIT_MQ"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	
	return &config
}
