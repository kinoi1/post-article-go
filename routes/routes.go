package routes

import (
	"go-post-article/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Article backend API Running",
		})
	})

	articleController := &controllers.ArticleController{DB: db}
	api := r.Group("/api")
	{
		api.POST("/article", articleController.CreatePost)

		api.GET("/article", articleController.GetPosts)

		api.GET("/article/:id", articleController.GetPostByID)

		api.PUT("/article/:id", articleController.UpdatePost)

		api.DELETE("/article/:id", articleController.DeletePost)
	}
}
