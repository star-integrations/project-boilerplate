package config

import (
	"context"

	"github.com/heetch/confita"
	"golang.org/x/xerrors"
)

// Config - config base for reading from outside
type Config struct {
	CORSAllowOrigins []string `config:"CORS_ALLOW_ORIGINS"` // CORS Allow Origins
	// TODO: add environment variable
}

var defaultConfig = Config{}

func readConfig(ctx context.Context) (*Config, error) {
	loader := confita.NewLoader()

	cfg := defaultConfig

	if err := loader.Load(ctx, &cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}

// ReadConfig - read config from environment variables
func ReadConfig(ctx context.Context) (*Config, error) {
	cfg, err := readConfig(ctx)

	if err != nil {
		return nil, xerrors.Errorf("failed to read config from env: %w", err)
	}

	return cfg, nil
}
