package inmemory_test

import (
	"errors"
	"testing"
	"uacademy/article/models"
	"uacademy/article/storage/inmemory"
)

func TestAddArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	err = IM.AddAuthor("626f1e10-58a2-414e-83c5-899b92ea0ff5", models.CreateAuthorModel{
		Firstname: "John",
		Lastname:  "Doe",
	})

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	err = IM.AddArticle("20455551-7263-4009-91bd-3fa6a10e3827", models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Impsume",
		},
		AuthorID: "626f1e10-58a2-414e-83c5-899b92ea0ff5",
	})

	if err != nil {
		t.Errorf("IM.AddArticle() got error: %v", err)
	}

	article, err := IM.GetArticleByID("20455551-7263-4009-91bd-3fa6a10e3827")
	if err != nil {
		t.Errorf("IM.AddArticle() got error: %v", err)
	}

	if article.Title != "Lorem" || article.Body != "Impsume" {
		t.Errorf("mitmatch between data")
	}

	err = IM.AddArticle("20455551-7263-4009-91bd-3fa6a10e3827", models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Impsume",
		},
		AuthorID: "408bdc75-65e8-4f72-8c65-e33855cfccbe",
	})

	expectedError := errors.New("author not found")
	if err == nil {
		t.Errorf("IM.AddArticle() expected error but got nil")
	} else {
		if err.Error() != expectedError.Error() {
			t.Errorf("IM.AddArticle() expected: %v, but got error: %v", expectedError, err)
		}
	}

	t.Log("Test has been finished")
}
