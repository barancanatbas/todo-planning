package mariadb

type MariadbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port" default:"3306"`
	DbName   string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	CharSet  string `mapstructure:"char-set" default:"utf8mb4"`
	Debug    bool   `mapstructure:"debug" default:"false"`
}
