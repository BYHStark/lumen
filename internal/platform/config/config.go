package config

import (
	"errors"
	"os"
)

func Load() error {
	dsn := os.Getenv("LUMEN_POSTGRES_DSN")
	if dsn == "" {
		return errors.New("LUMEN_POSTGRES_DSN 不能为空")
	}
	return nil
}
