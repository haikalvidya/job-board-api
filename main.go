package main

import (
	"flag"
	"job-board-api/cmd"
	"job-board-api/internal/routers"
	"job-board-api/migrations"
	"log"
)

func main() {
	configFile := flag.String("config", "config.yaml", "config file")
	migrate := flag.Bool("migrate", false, "migrate database")
	flag.Parse()
	cmd.Load(*configFile)
	if *migrate {
		migrations.Migrate()
	} else {
		routers.LoadRoutes(cmd.Http.Server.App)
		log.Println("Server started on port", cmd.Http.Server.Port)
		cmd.Http.Server.Serve()
	}
}
