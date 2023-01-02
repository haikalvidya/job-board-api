package config

import (
	"fmt"
	"time"

	"github.com/gofiber/session/v2"
	"github.com/gofiber/session/v2/provider/redis"
)

// session config for session on fiber

type SessionConfig struct {
	*session.Session
	Driver   string `mapstructure:"driver" yaml:"driver"`
	Name     string `mapstructure:"name" yaml:"name"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	DB       int    `mapstructure:"db" yaml:"db"`
	Password string `mapstructure:"password" yaml:"password"`
}

func (s *SessionConfig) SetUp() error {
	provider, err := redis.New(redis.Config{
		KeyPrefix:   "verify_rest_",
		Addr:        fmt.Sprint(s.Host, ":", s.Port),
		Password:    s.Password,
		DB:          s.DB,
		PoolSize:    8,
		IdleTimeout: 30 * time.Second,
	})
	if err != nil {
		return err
	}
	s.Session = session.New(session.Config{
		Provider: provider,
	})
	return nil
}
