package utils

import (
	config "github.com/furqonzt99/snackbox/configs"
	"github.com/furqonzt99/snackbox/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {

	conn := config.Database.Username + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Name + "?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitialMigrate(db *gorm.DB) {
	if config.Mode == "development" {
		db.Migrator().DropTable(&models.User{})
		db.Migrator().DropTable(&models.Product{})
		db.Migrator().DropTable(&models.Transaction{})
		db.Migrator().DropTable(&models.Partner{})
		db.Migrator().DropTable(&models.Rating{})
		db.Migrator().DropTable(&models.Cashout{})

		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Product{})
		db.AutoMigrate(&models.Transaction{})
		db.AutoMigrate(&models.Partner{})
		db.AutoMigrate(&models.Rating{})
		db.AutoMigrate(&models.Cashout{})
	} else {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Product{})
		db.AutoMigrate(&models.Transaction{})
		db.AutoMigrate(&models.Partner{})
		db.AutoMigrate(&models.Rating{})
		db.AutoMigrate(&models.Cashout{})
	}

}