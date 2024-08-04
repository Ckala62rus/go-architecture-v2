package utils

import (
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"practice/domains"
)

func AutoMigrateInitialize(db *gorm.DB, log *slog.Logger) {
	// initialize auto migration
	for _, models := range RegisterModel() {
		err := db.Debug().AutoMigrate(models.Model)

		if err != nil {
			log.Info(err.Error())
		}
	}

	fmt.Println("Database migrated successfully!")
}

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: domains.User{}},
	}
}
