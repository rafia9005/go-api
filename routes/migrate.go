package routes

import (
	"github.com/rafia9005/go-api/database"
)

func RunMigrate(dataModel interface{}) {
	database.DB.Debug().AutoMigrate(dataModel)
}
