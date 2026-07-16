package controllers

import (
	"go-post-article/models"
	"go-post-article/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleController struct {
	DB *gorm.DB
}

func NewArticleController(db *gorm.DB) *ArticleController {
	return &ArticleController{DB: db}
}

func (pc *ArticleController) CreatePost(c *gin.Context) {
	var p models.Posts

	if err := c.ShouldBindJSON(&p); err != nil {
		if errorMessages, isValidationError := utils.FormatValidationError(err); isValidationError {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload JSON tidak valid"})
		return
	}
	result := pc.DB.Create(&p)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (pc *ArticleController) GetPosts(c *gin.Context) {
	var posts []models.Posts

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, errLimit := strconv.Atoi(limitStr)
	offset, errOffset := strconv.Atoi(offsetStr)

	if errLimit != nil || errOffset != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter limit atau offset harus berupa angka"})
		return
	}

	result := pc.DB.Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *ArticleController) GetPostByID(c *gin.Context) {
	var post models.Posts
	id := c.Param("id")

	result := pc.DB.First(&post, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *ArticleController) UpdatePost(c *gin.Context) {
	var post models.Posts
	id := c.Param("id")
	if err := pc.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input models.Posts
	if err := c.ShouldBindJSON(&input); err != nil {
		if errorMessages, isValidationError := utils.FormatValidationError(err); isValidationError {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload JSON tidak valid"})
		return
	}

	if err := pc.DB.Model(&post).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (pc *ArticleController) DeletePost(c *gin.Context) {
	var post models.Posts
	id := c.Param("id")

	if err := pc.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
