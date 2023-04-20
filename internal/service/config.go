package service

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	Environment Environment `required:"true"`
	Port        uint        `required:"true"`
	RateLimit   string      `default:"1000-H"` // default to 1000 requests per hour
	DSN         string      `required:"true"`
}

func LoadConfig(logger *zap.Logger, envKey string, prefix string, conf *Config) error {
	if envFileName := os.Getenv(envKey); envFileName != "" {
		if err := godotenv.Load(envFileName); err != nil {
			logger.Error("error loading godotenv", zap.Error(err))
			return err
		}
	}

	err := envconfig.Process(prefix, conf)
	if err != nil {
		logger.Error("failed to parse config", zap.Error(err))
		return err
	}

	conf.Compute()

	return nil
}

func (c *Config) Compute() {
	if c.Environment == EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}
}
