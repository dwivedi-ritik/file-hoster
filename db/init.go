package db

import (
	"errors"

	"github.com/dwivedi-ritik/filehost-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddRow(product *models.Product) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return errors.New("can not connect with the database")
	}

	db.Create(&product)
	return nil
}

func GetRow(product models.Product) (models.Product, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	var userProduct models.Product

	if err != nil {
		return userProduct, errors.New("can not connect with the database")
	}
	db.First(&userProduct, product)

	if userProduct.Price == 0 {
		return userProduct, errors.New("row cannot be found")
	}

	return userProduct, nil
}

func GetAllRow() ([]models.Product, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	var userProduct []models.Product

	if err != nil {
		return userProduct, errors.New("can not connect with the database")
	}
	db.Find(&userProduct)

	return userProduct, nil

}

func MakeMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Product{})
}
