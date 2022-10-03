package handlers

import (
	"net/http"
	"time"

	"uacademy/article/models"
	"uacademy/article/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}

// CreateArticle godoc
// @Summary     Create article
// @Description create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModel true "article body"
// @Success     201     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/article [post]
func CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := storage.AddArticle(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := storage.GetArticleByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | GetList",
		Data:    article,
	})
}

// GetArticleByID godoc
// @Summary     get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Article}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/article/{id} [get]
func GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	article, err := storage.GetArticleByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=[]models.Article}
// @Router      /v2/article [get]
func GetArticleList(c *gin.Context) {
	articleList, err := storage.GetArticleList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    articleList,
	})
}

// UpdateArticle ...
func UpdateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, v := range storage.InMemoryArticleData {
		if v.ID == article.ID {
			article.CreatedAt = v.CreatedAt
			t := time.Now()
			article.UpdatedAt = &t

			storage.InMemoryArticleData[i] = article

			c.JSON(http.StatusOK, gin.H{
				"data":    storage.InMemoryArticleData,
				"message": "Article | Update",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article | Update | NOT FOUND",
		"data":    storage.InMemoryArticleData,
	})
}

// DeleteArticle ...
func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	for i, v := range storage.InMemoryArticleData {
		if v.ID == idStr {
			storage.InMemoryArticleData = remove(storage.InMemoryArticleData, i)
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
