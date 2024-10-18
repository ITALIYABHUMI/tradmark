package model

type DatabaseConfig struct {
	Host     string `json:"host" validate:"required" config:"DB_HOST"`
	Port     string `json:"port" validate:"required" config:"DB_PORT"`
	User     string `json:"user" validate:"required" config:"DB_USER"`
	Password string `json:"password" validate:"required" config:"DB_PASSWORD"`
	DBName   string `json:"db_name" validate:"required" config:"DB_NAME"`
	SSLMode  string `json:"ssl_mode" validate:"required" config:"DB_SSL_MODE"`
}

type ServerConfig struct {
	Port int `json:"port" validate:"required" config:"PORT"`
}

type Settings struct {
	DefaultConfig ServerConfig
	DBConfig      DatabaseConfig
}
