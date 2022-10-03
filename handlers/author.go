package handlers

import (
	"net/http"
	"uacademy/article/models"
	"uacademy/article/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "author body"
// @Success     201    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v2/author [post]
func CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	id := uuid.New()

	err := storage.AddAuthor(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := storage.GetAuthorByID(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | GetList",
		Data:    author,
	})
}

// GetAuthorByID godoc
// @Summary     get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Param       id path string true "Author ID"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/author/{id} [get]
func GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation

	author, err := storage.GetAuthorByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthorList godoc
// @Summary     List author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=[]models.Author}
// @Router      /v2/author [get]
func GetAuthorList(c *gin.Context) {
	authorList, err := storage.GetAuthorList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    authorList,
	})
}
