package database

import (
	"log"
	"os"
	"voting-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Inisialisasi database
func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("root@tcp(127.0.0.1:3306)/my_app")
	if dsn == "" {
		log.Fatal("DB_DSN tidak boleh kosong")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database", err)
	}

	// Melakukan migrasi model ke database
	if err := db.AutoMigrate(&models.Candidate{}, &models.Voter{}); err != nil {
		log.Fatal("Gagal migrasi tabel", err)
	}

	return db
}
