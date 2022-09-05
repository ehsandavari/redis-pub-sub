package config

import (
	DomainEnums "OrderSubscriber/Domain/Enums"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type SConfiguration struct {
	MySql struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"mysql"`
	Redis struct {
		URL    string                                              `mapstructure:"url"`
		Queues map[DomainEnums.RedisQueues]DomainEnums.RedisQueues `mapstructure:"queues"`
	} `mapstructure:"redis"`
}

func NewConfiguration() SConfiguration {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config.dev" + os.Getenv("env"))
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
