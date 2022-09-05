package config

import (
	DomainEnums "OrderPublisher/Domain/Enums"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type SConfiguration struct {
	Application struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"application"`
	Redis struct {
		URL    string                                              `mapstructure:"url"`
		Queues map[DomainEnums.RedisQueues]DomainEnums.RedisQueues `mapstructure:"queues"`
	} `mapstructure:"redis"`
}

func NewConfiguration() SConfiguration {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config." + os.Getenv("env"))
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err, "Fatal error reading config")
	} else {
		log.WithFields(log.Fields{"config": viper.ConfigFileUsed()}).Info("Viper config")
	}

	configuration := SConfiguration{}
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config File")
	}
	return configuration
}
