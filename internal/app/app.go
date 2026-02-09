package app

import (
	"github.com/scmbr/auth-service/internal/app/config"
	"github.com/scmbr/auth-service/pkg/logger"
)

func Run(configsDir string) {
	logger.Init()
	cfg, err := config.Init(configsDir)
	if err != nil {
		logger.Error("failed to initialize configs", err, nil)
	}
	logger.Info("configs initialized successfully", map[string]interface{}{
		"environment": cfg.Environment,
	})
}
