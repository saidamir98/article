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
	article.Author = entity.Author
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func GetArticleByID(id string) (models.Article, error) {
	for _, v := range InMemoryArticleData {
		if v.ID == id {
			return v, nil
		}
	}
	return models.Article{}, errors.New("article not found")
}

// GetArticleList ...
func GetArticleList() (resp []models.Article, err error) {
	resp = InMemoryArticleData
	return resp, err
}
