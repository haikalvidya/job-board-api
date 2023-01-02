package cmd

import (
	"job-board-api/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var Http *config.AppConfig

var Version = "develop"

func Load(configFile string) {
	Http = &config.AppConfig{ConfigFile: configFile}
	err := Http.SetUp()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	LoadBuiltInMiddlewares(Http)
}

func LoadBuiltInMiddlewares(app *config.AppConfig) {
	app.Server.Use(recover.New())
	app.Server.Use(etag.New())
	app.Server.Use(compress.New(compress.Config{
		Level: 1,
	}))
	if app.Server.Debug {
		app.Server.Use(pprof.New())
	}
}
