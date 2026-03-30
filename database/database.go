package database

import (
	"log"
	"os"
	"voting-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Baca DSN dari environment variable DB_DSN
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN tidak boleh kosong")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database: ", err)
	}

	// Migrasi model
	if err := db.AutoMigrate(&models.Candidate{}, &models.Voter{}); err != nil {
		log.Fatal("Gagal migrasi tabel: ", err)
	}

	return db
}