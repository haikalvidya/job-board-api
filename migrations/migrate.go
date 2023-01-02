package migrations

import (
	"job-board-api/cmd"
	"log"
)

func Migrate() {
	err := cmd.Http.Database.Migrator().AutoMigrate()
	if err != nil {
		panic(err)
	}
	log.Println("Migrations completed successfully")
}
