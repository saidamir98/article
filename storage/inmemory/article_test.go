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

	errorAuthorNotFound := errors.New("author not found")

	authorID := "626f1e10-58a2-414e-83c5-899b92ea0ff5"
	authorData := models.CreateAuthorModel{
		Firstname: "John",
		Lastname:  "Doe",
	}
	notFoundAuthorID := "63f0307d-8fa9-474f-a438-77319effc9ca"
	content := models.Content{
		Title: "Lorem",
		Body:  "Impsume",
	}

	err = IM.AddAuthor(authorID, authorData)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	var tests = []struct {
		name       string
		id         string
		data       models.CreateArticleModel
		wantError  error
		wantResult models.PackedArticleModel
	}{
		{
			name: "success",
			id:   "20455551-7263-4009-91bd-3fa6a10e3827",
			data: models.CreateArticleModel{
				Content:  content,
				AuthorID: authorID,
			},
			wantError: nil,
			wantResult: models.PackedArticleModel{
				Content: content,
			},
		},
		{
			name: "fail",
			id:   "30455551-7263-4009-91bd-3fa6a10e3827",
			data: models.CreateArticleModel{
				Content:  content,
				AuthorID: notFoundAuthorID,
			},
			wantError:  errorAuthorNotFound,
			wantResult: models.PackedArticleModel{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = IM.AddArticle(tt.id, tt.data)

			if tt.wantError == nil {
				if err != nil {
					t.Errorf("IM.AddArticle() got error: %v", err)
				}

				article, err := IM.GetArticleByID(tt.id)
				if err != nil {
					t.Errorf("IM.AddArticle() got error: %v", err)
				}

				if tt.wantResult.Content != article.Content {
					t.Errorf("IM.AddArticle() expected: %v but got: %v", tt.wantResult.Content, article.Content)
				}
			} else {
				if tt.wantError.Error() != err.Error() {
					t.Errorf("IM.AddArticle() expected error: %v, but got error: %v", tt.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}

func TestGetArticleByID(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	errorArticleNotFound := errors.New("article not found")

	authorID := "626f1e10-58a2-414e-83c5-899b92ea0ff5"
	authorData := models.CreateAuthorModel{
		Firstname: "John",
		Lastname:  "Doe",
	}
	err = IM.AddAuthor(authorID, authorData)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	articleID := "20455551-7263-4009-91bd-3fa6a10e3827"
	articleData := models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Impsume",
		},
		AuthorID: authorID,
	}
	err = IM.AddArticle(articleID, articleData)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	var tests = []struct {
		name       string
		id         string
		mockFunc   func()
		wantError  error
		wantResult models.PackedArticleModel
	}{
		{
			name: "success",
			id:   articleID,
			mockFunc: func() {

			},
			wantError: nil,
			wantResult: models.PackedArticleModel{
				Content: articleData.Content,
			},
		},
		{
			name: "fail: artilce not found",
			id:   "30455551-7263-4009-91bd-3fa6a10e3827",
			mockFunc: func() {

			},
			wantError:  errorArticleNotFound,
			wantResult: models.PackedArticleModel{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article, err := IM.GetArticleByID(tt.id)
			if tt.wantError == nil {
				if err != nil {
					t.Errorf("IM.GetArticleByID() got error: %v", err)
				}

				if tt.wantResult.Content != article.Content {
					t.Errorf("IM.GetArticleByID() expected: %v but got: %v", tt.wantResult.Content, article.Content)
				}
			} else {
				if tt.wantError.Error() != err.Error() {
					t.Errorf("IM.GetArticleByID() expected error: %v, but got error: %v", tt.wantError, err)
				}
			}
		})
	}
	t.Log("Test has been finished")
}
