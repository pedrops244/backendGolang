package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "mysql://root:yZuxhlnVDFcDsQELWJBbFAFPPaXaspTw@autorack.proxy.rlwy.net:36112/railway."),
		Port:       getEnv("PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "yZuxhlnVDFcDsQELWJBbFAFPPaXaspTw"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "mysql.railway.internal"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "railway"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
