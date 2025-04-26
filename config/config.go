package config

import (
	"os"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		Port:      getEnv("PORT", "8081"),
		DBHost:    getEnv("DB_HOST", "db"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "test"),
		DBPass:    getEnv("DB_PASS", "test"),
		DBName:    getEnv("DB_NAME", "test"),
		JWTSecret: getEnv("JWT_SECRET", "your_jwt_secret"),
	}
}

func InitDB(cfg *Config) *gorm.DB {
	dsn := "host=" + cfg.DBHost + 
		" user=" + cfg.DBUser + 
		" password=" + cfg.DBPass + 
		" dbname=" + cfg.DBName + 
		" port=" + cfg.DBPort + 
		" sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}