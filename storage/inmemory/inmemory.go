package inmemory

import "uacademy/article/models"

// InMemory ...
type InMemory struct {
	Db *DB
}

// DB mock
type DB struct {
	// InMemoryArticleData ...
	InMemoryArticleData []models.Article
	// InMemoryAuthorData ...
	InMemoryAuthorData []models.Author
}
