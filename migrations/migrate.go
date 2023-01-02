package migrations

import (
	"job-board-api/cmd"
	"job-board-api/internal/models"
	"log"
)

func Migrate() {
	err := cmd.Http.Database.Migrator().AutoMigrate(
		&models.User{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migrations completed successfully")
}
