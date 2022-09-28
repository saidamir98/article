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

// CreateArticle godoc
// @Summary      Create article
// @Description  create a new article
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param article body Article true "article body"
// @Success      201  {object}  JSONResponse{data=[]Article}
// @Failure      400  {object}  JSONErrorResponse
// @Router       /v2/article [post]
func CreateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	article.ID = id.String()
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	c.JSON(http.StatusCreated, JSONResponse{
		Message: "Article | GetList",
		Data:    InMemoryArticleData,
	})
}

// GetArticleByID godoc
// @Summary      get article by id
// @Description  get an article by id
// @Tags         articles
// @Accept       json
// @Param        id   path      string  true  "Article ID"
// @Produce      json
// @Success      200  {object}  JSONResponse{data=Article}
// @Failure      400  {object}  JSONErrorResponse
// @Router       /v2/article/{id} [get]
func GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	for _, v := range InMemoryArticleData {
		if v.ID == idStr {
			c.JSON(http.StatusOK, JSONResponse{
				Message: "Article | GetByID",
				Data:    v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, JSONErrorResponse{
		Error: "Article | GetByID | NOT FOUND",
	})
}

// GetArticleList godoc
// @Summary      List articles
// @Description  get articles
// @Tags         articles
// @Accept       json
// @Produce      json
// @Success      200  {object}   JSONResponse{data=[]Article}
// @Router       /v2/article [get]
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, JSONResponse{
		Message: "Article | GetList",
		Data:    InMemoryArticleData,
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
