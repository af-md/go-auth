package model

type DBConfig struct {
	DbName     string `env:"MYSQL_DB"`
	Username   string `env:"MYSQL_USER"`
	Password   string `env:"MYSQL_PASSWORD"`
	Connection string `env:"MYSQL_CONNECTION"`
	Host       string `env:"MYSQL_HOST"`
	Port       string `env:"MYSQL_PORT"`
}
