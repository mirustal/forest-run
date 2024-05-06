package jwt

import (
	"errors"
	"time"
)

type JWTConfig struct {
	SecureKey            string        `yaml:"secureKey"`
	JWTTokenLifeTime     time.Duration `yaml:"JWTTokenLifeTime"`
	RefreshTokenLifeTime time.Duration `yaml:"refreshTokenLifeTime"`
}

func (cfg JWTConfig) Validate() error {
	if cfg.JWTTokenLifeTime <= 0 {
		return errors.New("JWTTokenLifeTime <= 0")
	}

	if len(cfg.SecureKey) == 0 {
		return errors.New("length of secure key is 0")
	}

	if cfg.RefreshTokenLifeTime <= 0 {
		return errors.New("RefreshTokenLifeTime <= 0")
	}

	return nil
}
