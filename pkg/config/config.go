package config

import (
	"strings"

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
	Protocal string `mapstructure:"PROTOCOL"`
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	VHost string `mapstructure:"VHOST"`
}

type Config struct {
	Port string `mapstructure:"PORT"`
	DB DB `mapstructure:"DB"`
	RabbitMQ RabbitMQ `mapstructure:"RABBITMQ"`
}

func LoadConfig() *Config {
	var config Config

	replacer := strings.NewReplacer(".", "_")
    	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.AutomaticEnv()

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
