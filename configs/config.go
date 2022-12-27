package configs

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		db, err := InitDB()
		if err != nil {
			panic("failed to connect database")
		}

		config = &Config{DB: db}
	}

	return config
}

func (config Config) ContextDB(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Set("db", config.DB)
		return next(context)
	}
}
