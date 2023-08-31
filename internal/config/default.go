package config

import (
	"github.com/khulnasoft-lab/oss-chain-bench/internal/logger"
)

func loadDefaultConfiguration() *Configuration {
	return &Configuration{
		LogConfiguration: &LogConfiguration{
			LogLevel:  "info",
			LogFormat: logger.NormalFormat,
		},
	}
}
