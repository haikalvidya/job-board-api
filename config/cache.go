package config

import "github.com/gofiber/storage/redis"

// using redis
type CacheConfig struct {
	*redis.Storage
	Driver   string `mapstructure:"driver" yaml:"driver"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	DB       int    `mapstructure:"db" yaml:"db"`
	Password string `mapstructure:"password" yaml:"password"`
}

func (c *CacheConfig) SetUp() {
	c.Storage = redis.New(redis.Config{
		Host:     c.Host,
		Port:     c.Port,
		Password: c.Password,
		Database: c.DB,
	})
}
