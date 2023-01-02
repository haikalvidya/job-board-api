package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Server     ServerConfig   `mapstructure:"server" yaml:"server" env:"server"`
	Database   DatabaseConfig `mapstructure:"database" yaml:"database" env:"database"`
	Jwt        JwtConfig      `mapstructure:"jwt" yaml:"jwt" env:"jwt"`
	Session    SessionConfig  `mapstructure:"session" yaml:"session" env:"session"`
	Cache      CacheConfig    `mapstructure:"cache" yaml:"cache" env:"cache"`
	ConfigFile string
	Hash       Hash
}

func (cfg *AppConfig) SetUp() error {
	if err := cleanenv.ReadConfig(cfg.ConfigFile, cfg); err != nil {
		return err
	}
	cfg.Server.SetUp()
	if err := cfg.Database.SetUp(); err != nil {
		return err
	}
	if err := cfg.Session.SetUp(); err != nil {
		return err
	}
	cfg.Cache.SetUp()
	return nil
}
