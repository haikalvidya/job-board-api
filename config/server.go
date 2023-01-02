package config

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"emperror.dev/errors"
	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	*fiber.App
	Name  string `mapstructure:"name" yaml:"name" env:"name"`
	Host  string `mapstructure:"host" yaml:"host" env:"host"`
	Port  string `mapstructure:"port" yaml:"port" env:"port"`
	Debug bool   `mapstructure:"debug" yaml:"debug" env:"debug" default:"true"`
	Url   string `mapstructure:"url" yaml:"url" env:"url" default:"http://localhost"`
}

func (s *ServerConfig) SetUp() {
	s.App = fiber.New(fiber.Config{
		ServerHeader:          s.Name,
		DisableStartupMessage: true,
		ReduceMemoryUsage:     true,
		ErrorHandler:          CustomErrorHandler,
		Concurrency:           256 * 1024 * 1024,
	})
}

func (s *ServerConfig) Serve() error {
	a := s.Host + ":" + s.Port
	go func() {
		if err := s.Listen(a); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGABRT,
		syscall.SIGQUIT,
	)
	<-c
	fmt.Println("I'm shutting down")
	return s.Shutdown()
}

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	// StatusCode defaults to 500
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	er := errors.WithStack(err)
	fmt.Printf("%+v", er)
	fmt.Printf("%+v", err)
	return c.Status(code).JSON(err)
}
