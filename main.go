package main

import (
	"go-post-article/config"
	"go-post-article/routes"
	"go-post-article/seeders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.MigrateDB()
	seeders.SeedPosts()
	r := gin.Default()

	r.Use(cors.New(config.CORSConfig()))

	routes.SetupRoutes(r, config.DB)

	r.Run(":8080")
}
