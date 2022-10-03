package storage

import (
	"errors"
	"time"
	"uacademy/article/models"
)

// InMemoryAuthorData ...
var InMemoryAuthorData []models.Author

// AddAuthor ...
func AddAuthor(id string, entity models.CreateAuthorModel) error {
	var author models.Author
	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.CreatedAt = time.Now()

	InMemoryAuthorData = append(InMemoryAuthorData, author)
	return nil
}

// GetAuthorByID ...
func GetAuthorByID(id string) (models.Author, error) {
	var result models.Author
	for _, v := range InMemoryAuthorData {
		if v.ID == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

// GetAuthorList ...
func GetAuthorList() (resp []models.Author, err error) {
	resp = InMemoryAuthorData
	return resp, err
}
