package sqlite

import (
	"basic-rest/pkg/model"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.Migrator().DropTable(&model.Product{})
	DB.AutoMigrate(&model.Product{})

	p := model.Product{
		Name:        "Basic computer",
		Provider:    "Prov Inc.",
		Quantity:    20,
		Price:       200.00,
		Description: "Basic product.",
	}
	DB.Create(&p)
}
