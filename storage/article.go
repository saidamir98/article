package storage

import (
	"errors"
	"strings"
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
	// Check author
	article.AuthorID = entity.AuthorID
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.ID == id && v.DeletedAt == nil {
			author, err := GetAuthorByID(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
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
func GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {
	off := 0
	c := 0
	for _, v := range InMemoryArticleData {
		if v.DeletedAt == nil && (strings.Contains(v.Title, search) || strings.Contains(v.Body, search)) {
			if offset <= off {
				c++
				resp = append(resp, v)
			}

			if c >= limit {
				break
			}

			off++
		}
	}

	return resp, err
}

// UpdateArticle ...
func UpdateArticle(entity models.UpdateArticleModel) error {
	for i, v := range InMemoryArticleData {
		if v.ID == entity.ID && v.DeletedAt == nil {
			v.Content = entity.Content
			t := time.Now()
			v.UpdatedAt = &t
			InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}

// DeleteArticle ...
func DeleteArticle(id string) error {
	for i, v := range InMemoryArticleData {
		if v.ID == id {
			if v.DeletedAt != nil {
				return errors.New("article already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			InMemoryArticleData[i] = v
			return nil
		}
	}

	return errors.New("article not found")
}
