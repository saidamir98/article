package storage

import (
	"errors"
	"time"
	"uacademy/article/models"
)

// InMemoryArticleData ...
var InMemoryArticleData []models.Article

// AddArticle ...
func AddArticle(id string, entity models.CreateArticleModel) error {
	var article models.Article
	article.ID = id
	article.Content = entity.Content
	article.AuthorID = entity.AuthorID
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.ID == id {
			author, err := GetAuthorByID(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.Content = v.Content
			result.Author = author
			result.CreatedAt = v.CreatedAt
			result.UpdatedAt = v.UpdatedAt
			result.DeletedAt = v.DeletedAt
			return result, nil
		}
	}
	return result, errors.New("article not found")
}

// GetArticleList ...
func GetArticleList() (resp []models.Article, err error) {
	resp = InMemoryArticleData
	return resp, err
}
