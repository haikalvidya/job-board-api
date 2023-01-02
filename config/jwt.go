package config

type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}
