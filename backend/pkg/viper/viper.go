package viper

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupConfig(conf interface{}, args ...string) {
	if len(args) == 0 {
		args = append(args, "config")
	}

	if len(args) != 2 {
		args = append(args, ".")
	}

	viper.SetConfigName(args[0])
	viper.SetConfigType("yaml")
	viper.AddConfigPath(args[1])

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msgf("error to reading config file")
	}

	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msgf("error to decode")
	}
}
