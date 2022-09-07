package config

import (
	DomainEnums "OrderSubscriber/Domain/Enums"
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

func NewConfiguration() (SConfiguration, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config.dev" + os.Getenv("env"))
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return SConfiguration{}, err
	}

	configuration := SConfiguration{}
	if err := viper.Unmarshal(&configuration); err != nil {
		return SConfiguration{}, err
	}
	return configuration, nil
}
