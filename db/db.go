package db

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/dwivedi-ritik/filehost-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB_PATH string = dbPath()
)

func dbPath() string {
	return filepath.Join(os.Getenv("HOME"), ".config/filehost/database.db")
}

func AddRow(file *models.File) (uint, error) {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		return 0, errors.New("can not connect with the database")
	}

	db.Create(&file)
	return file.ID, nil
}

func GetRow(File models.File) (models.File, error) {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	var userFile models.File

	if err != nil {
		return userFile, errors.New("can not connect with the database")
	}
	db.First(&userFile, File)

	if userFile.ID == 0 {
		return userFile, errors.New("row cannot be found")
	}

	return userFile, nil
}

func GetAllRow() ([]models.File, error) {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	var userFile []models.File

	if err != nil {
		return userFile, errors.New("can not connect with the database")
	}
	db.Find(&userFile)

	return userFile, nil

}

func MakeMigration() {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.File{})
}
