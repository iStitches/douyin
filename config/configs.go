package config

type ServerConfig struct {
	Name      string      `mapstructure:"name"`
	Port      int         `mapstructure:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
	JWTKey    JWTConfig   `mapstructure:"jwt"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
