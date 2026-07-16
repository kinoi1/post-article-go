package routes

import (
	"go-post-article/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Tambahkan parameter db *gorm.DB agar bisa dioper ke controller
func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Absensi API Running",
		})
	})

	// Inisialisasi controller di sini
	articleController := &controllers.ArticleController{DB: db}
	api := r.Group("/api")
	{
		// 1. Membuat article baru
		api.POST("/article", articleController.CreatePost)

		// 2. Menampilkan seluruh article dengan paging (limit & offset)
		api.GET("/article", articleController.GetPosts)

		// 3. Menampilkan article berdasarkan id
		api.GET("/article/:id", articleController.GetPostByID)

		// 4. Merubah data article berdasarkan id
		api.PUT("/article/:id", articleController.UpdatePost)

		// 5. Menghapus data article berdasarkan id
		api.DELETE("/article/:id", articleController.DeletePost)
	}
}
