package settings

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Settings struct {
	DatabaseUrl string `env:"DATABASE_URL"`
	Port        int    `env:"PORT,default=8080"`
	SecretKey   string `env:"SECRET_KEY"`
}

var globalSettings Settings

func LoadConfig(ctx context.Context) error {
	return envconfig.Process(ctx, &globalSettings)
}

func GetSettings() Settings {
	return globalSettings
}
