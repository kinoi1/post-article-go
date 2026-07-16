package seeders

import (
	"time"

	"go-post-article/config"
	"go-post-article/models"
)

func SeedPosts() error {
	var count int64
	config.DB.Model(&models.Posts{}).Count(&count)

	if count > 0 {
		return nil
	}
	posts := []models.Posts{
		{
			Title:       "Belajar Golang untuk Pemula",
			Content:     "Panduan lengkap memulai pemrograman menggunakan Golang.",
			Category:    "Programming",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Mengenal Framework Gin",
			Content:     "Framework Gin merupakan salah satu framework web tercepat di Golang.",
			Category:    "Programming",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Cara Menggunakan GORM",
			Content:     "Tutorial penggunaan GORM sebagai ORM untuk Golang.",
			Category:    "Database",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Tips Optimasi Query SQL",
			Content:     "Beberapa tips untuk meningkatkan performa query database.",
			Category:    "Database",
			Status:      "Draft",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Mengenal REST API",
			Content:     "Konsep dasar REST API dan implementasinya.",
			Category:    "Backend",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "JWT Authentication dengan Golang",
			Content:     "Implementasi autentikasi menggunakan JWT.",
			Category:    "Security",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Membuat CRUD dengan Gin dan GORM",
			Content:     "Tutorial membuat aplikasi CRUD sederhana.",
			Category:    "Programming",
			Status:      "Draft",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Deploy Aplikasi Golang ke VPS",
			Content:     "Langkah-langkah deploy aplikasi Golang ke server Linux.",
			Category:    "DevOps",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Pengenalan Docker",
			Content:     "Memahami dasar-dasar container menggunakan Docker.",
			Category:    "DevOps",
			Status:      "Publish",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Title:       "Artikel Lama yang Dihapus",
			Content:     "Artikel ini sudah tidak digunakan lagi.",
			Category:    "Archive",
			Status:      "Thrash",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
	}

	return config.DB.Create(&posts).Error
}
