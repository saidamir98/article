package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// InMemoryArticleData ...
var InMemoryArticleData []Article

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}

// CreateArticle ...
func CreateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New()
	article.ID = id.String()
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	c.JSON(http.StatusCreated, gin.H{
		"data":    InMemoryArticleData,
		"message": "Article | Create",
	})
}

// GetArticleByID ...
var GetArticleByID = func(c *gin.Context) {
	idStr := c.Param("id")

	for _, v := range InMemoryArticleData {
		if v.ID == idStr {
			c.JSON(http.StatusOK, gin.H{
				"message": "Article | GetByID",
				"data":    v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article | GetByID | NOT FOUND",
		"data":    nil,
	})
}

// GetArticleList ...
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Article | GetList",
		"data":    InMemoryArticleData,
	})
}

// UpdateArticle ...
func UpdateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, v := range InMemoryArticleData {
		if v.ID == article.ID {
			article.CreatedAt = v.CreatedAt
			t := time.Now()
			article.UpdatedAt = &t

			InMemoryArticleData[i] = article

			c.JSON(http.StatusOK, gin.H{
				"data":    InMemoryArticleData,
				"message": "Article | Update",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article | Update | NOT FOUND",
		"data":    InMemoryArticleData,
	})
}

// DeleteArticle ...
func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	for i, v := range InMemoryArticleData {
		if v.ID == idStr {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			c.JSON(http.StatusOK, gin.H{
				"message": "Article | Delete",
				"data":    v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article | Delete | NOT FOUND",
		"data":    nil,
	})
}
