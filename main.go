package main

import (
	"flag"
	"job-board-api/cmd"
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
		// TODO add routes
		log.Fatal(cmd.Http.Server.Serve())
	}
}
