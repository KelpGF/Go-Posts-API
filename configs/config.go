package configs

import "os"

type ConfigType struct {
	HTTPHost   string
	HTTPPort   string
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func NewConfig() *ConfigType {
	return &ConfigType{
		HTTPHost:   os.Getenv("HTTP_HOST"),
		HTTPPort:   os.Getenv("HTTP_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

func (c *ConfigType) GetHttpHost() string {
	if c.HTTPHost == "" {
		return "localhost"
	}

	return c.HTTPHost
}

func (c *ConfigType) GetHttpPort() string {
	if c.HTTPPort == "" {
		return "3000"
	}

	return c.HTTPPort
}
