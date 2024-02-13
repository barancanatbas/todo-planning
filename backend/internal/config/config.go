package config

import "task-manager/pkg/mariadb"

type App struct {
	Port        string `mapstructure:"port" default:"80"`
	BasePath    string `mapstructure:"base-path" default:"/"`
	LogRequests bool   `mapstructure:"log-requests" default:"true"`
}

type LogConfig struct {
	Level string `mapstructure:"level" default:"info"`
}

type Config struct {
	App     App                   `mapstructure:"app"`
	MariaDB mariadb.MariadbConfig `mapstructure:"maria-db"`
	Log     LogConfig             `mapstructure:"log"`
}
