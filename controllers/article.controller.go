package controllers

import (
	"go-post-article/models"
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

// 1. POST /article - Membuat article baru
func (pc *ArticleController) CreatePost(c *gin.Context) {
	var p models.Posts

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload tidak valid"})
		return
	}

	if p.Status != "Publish" && p.Status != "Draft" && p.Status != "Thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus berupa 'Publish', 'Draft', atau 'Thrash'"})
		return
	}

	result := pc.DB.Create(&p)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Mengembalikan object kosong {} sesuai ketentuan
	c.JSON(http.StatusCreated, gin.H{})
}

// 2. GET /article - Menampilkan seluruh article dengan paging (limit & offset)
func (pc *ArticleController) GetPosts(c *gin.Context) {
	var posts []models.Posts

	// Mengambil query param limit dan offset, berikan nilai default jika kosong
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, errLimit := strconv.Atoi(limitStr)
	offset, errOffset := strconv.Atoi(offsetStr)

	if errLimit != nil || errOffset != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter limit atau offset harus berupa angka"})
		return
	}

	// Query ke database menggunakan Limit dan Offset dari GORM
	result := pc.DB.Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// 3. GET /article/:id - Menampilkan article berdasarkan ID
func (pc *ArticleController) GetPostByID(c *gin.Context) {
	var post models.Posts
	id := c.Param("id")

	// Mencari data berdasarkan ID
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

// 4. PUT /article/:id - Merubah data article berdasarkan ID
func (pc *ArticleController) UpdatePost(c *gin.Context) {
	var post models.Posts
	id := c.Param("id")

	// Cek apakah data artikelnya ada di database
	if err := pc.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Bind data baru dari request body
	var input models.Posts
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload tidak valid"})
		return
	}

	if input.Status != "" && input.Status != "Publish" && input.Status != "Draft" && input.Status != "Thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus berupa 'Publish', 'Draft', atau 'Thrash'"})
		return
	}

	// Update data menggunakan GORM Updates
	pc.DB.Model(&post).Updates(input)

	// Mengembalikan object kosong {} sesuai ketentuan
	c.JSON(http.StatusOK, gin.H{})
}

// 5. DELETE /article/:id - Menghapus data article berdasarkan ID
func (pc *ArticleController) DeletePost(c *gin.Context) {
	var post models.Posts
	id := c.Param("id")

	// Cek apakah data artikelnya ada
	if err := pc.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hapus data (Hard Delete)
	// Jika ingin soft delete, pastikan model.Post menggunakan gorm.DeletedAt
	if err := pc.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan object kosong {} atau status no content
	c.JSON(http.StatusOK, gin.H{})
}
