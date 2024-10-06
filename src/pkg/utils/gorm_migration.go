package utils

import (
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"practice/domains"
)

func AutoMigrateInitialize(db *gorm.DB, log *slog.Logger, env string) {
	// initialize auto migration
	for _, models := range RegisterModel() {
		var err error
		if env == "local" || env == "dev" {
			err = db.Debug().AutoMigrate(models.Model)
		} else {
			err = db.AutoMigrate(models.Model)
		}

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
