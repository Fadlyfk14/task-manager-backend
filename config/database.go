package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env (aman walau di Railway tidak ada)
	_ = godotenv.Load()

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("❌ SUPABASE_DSN tidak ditemukan. Pastikan environment variable sudah diset.")
	}

	// Pastikan sslmode=require (wajib untuk Supabase / Railway)
	if !strings.Contains(dsn, "sslmode=") {
		if strings.Contains(dsn, "?") {
			dsn += "&sslmode=require"
		} else {
			dsn += "?sslmode=require"
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal konek ke database: %v", err)
	}

	DB = db
	fmt.Println("✅ Koneksi ke PostgreSQL berhasil")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("❌ DB belum diinisialisasi. Panggil InitDB() dulu.")
	}
	return DB
}
