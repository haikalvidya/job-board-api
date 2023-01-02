package config

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DatabaseDriver struct {
	Driver   string `mapstructure:"driver" yaml:"driver" env:"driver"`
	Host     string `mapstructure:"host" yaml:"host" env:"host"`
	Port     string `mapstructure:"port" yaml:"port" env:"port"`
	Database string `mapstructure:"database" yaml:"database" env:"database"`
	Username string `mapstructure:"username" yaml:"username" env:"username"`
	Password string `mapstructure:"password" yaml:"password" env:"password"`
}

type DatabaseConfig struct {
	*gorm.DB
	Drivers        map[string]DatabaseDriver `yaml:"drivers"`
	DatabaseDriver DatabaseDriver            `mapstructure:"default" yaml:"default" env:"default"`
}

func (d *DatabaseConfig) SetUp() error {
	var err error
	var connectionString string
	if d.DB != nil {
		return nil
	}
	switch d.DatabaseDriver.Driver {
	case "mysql":
		connectionString = d.DatabaseDriver.Username + ":" + d.DatabaseDriver.Password + "@tcp(" + d.DatabaseDriver.Host + ":" + d.DatabaseDriver.Port + ")/" + d.DatabaseDriver.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
		d.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	case "postgres":
		connectionString = "host=" + d.DatabaseDriver.Host + " port=" + d.DatabaseDriver.Port + " user=" + d.DatabaseDriver.Username + " dbname=" + d.DatabaseDriver.Database + " password=" + d.DatabaseDriver.Password + " sslmode=disable"
		d.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	default:
		// default is mysql
		connectionString = d.DatabaseDriver.Username + ":" + d.DatabaseDriver.Password + "@tcp(" + d.DatabaseDriver.Host + ":" + d.DatabaseDriver.Port + ")/" + d.DatabaseDriver.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
		d.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	}

	if err != nil {
		return err
	}

	d.DB.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(100),
	)

	return nil
}
