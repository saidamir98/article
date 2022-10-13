package inmemory_test

import (
	"errors"
	"testing"
	"uacademy/article/models"
	"uacademy/article/storage/inmemory"
)

func InitData(stg *inmemory.InMemory) error {
	err := stg.AddAuthor("eb9eb7f9-47ca-4f75-815b-cd5c901b4df2", models.CreateAuthorModel{
		Firstname: "John",
		Lastname:  "Doe",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "1",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "2",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "3",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "4",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "5",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "1",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "2",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "3",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "4",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	err = stg.AddArticle("b6c46a35-d78b-4f1d-80fe-8d617f83ec6c", models.CreateArticleModel{
		Content: models.Content{
			Title: "5",
			Body:  "Impsume smth smth smthsmth",
		},
		AuthorID: "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2",
	})
	if err != nil {
		return err
	}

	return nil
}

func TestAddArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	InitData(&IM)

	errorAuthorNotFound := errors.New("author not found")

	authorID := "eb9eb7f9-47ca-4f75-815b-cd5c901b4df2"

	notFoundAuthorID := "63f0307d-8fa9-474f-a438-77319effc9ca"

	content := models.Content{
		Title: "Lorem",
		Body:  "Impsume",
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
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	InitData(&IM)

	errorArticleNotFound := errors.New("article not found")
	articleID := "b6c46a35-d78b-4f1d-80fe-8d617f83ec6c"

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
				Content: models.Content{
					Title: "Lorem",
					Body:  "Impsume smth smth smthsmth",
				},
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

func TestGetArticleList(t *testing.T) {
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	InitData(&IM)

	var tests = []struct {
		name           string
		offset         int
		limit          int
		search         string
		wantError      error
		numberOfResult int
	}{
		{
			name:           "success default",
			offset:         0,
			limit:          10,
			search:         "",
			wantError:      nil,
			numberOfResult: 10,
		},

		{
			name:           "success limit",
			offset:         0,
			limit:          5,
			search:         "",
			wantError:      nil,
			numberOfResult: 5,
		},
		{
			name:           "success offset",
			offset:         2,
			limit:          10,
			search:         "",
			wantError:      nil,
			numberOfResult: 9,
		},
		{
			name:           "success mix offset limit",
			offset:         2,
			limit:          5,
			search:         "",
			wantError:      nil,
			numberOfResult: 5,
		},
		{
			name:           "success out of offset",
			offset:         10,
			limit:          5,
			search:         "",
			wantError:      nil,
			numberOfResult: 1,
		},
		{
			name:           "success Lorem",
			offset:         0,
			limit:          10,
			search:         "Lorem",
			wantError:      nil,
			numberOfResult: 1,
		},
		{
			name:           "success Lorem",
			offset:         0,
			limit:          10,
			search:         "5",
			wantError:      nil,
			numberOfResult: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articleList, err := IM.GetArticleList(tt.offset, tt.limit, tt.search)

			if tt.wantError == nil {
				if err != nil {
					t.Errorf("IM.TestGetArticleList() got error: %v", err)
				}

				if tt.numberOfResult != len(articleList) {
					t.Errorf("IM.TestGetArticleList() expected: %d objects but got: %d objects", tt.numberOfResult, len(articleList))
				}
			} else {
				if tt.wantError.Error() != err.Error() {
					t.Errorf("IM.TestGetArticleList() expected error: %v, but got error: %v", tt.wantError, err)
				}
			}
		})
	}

	t.Log("Test has been finished")
}

func TestUpdateArticle(t *testing.T) {
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	InitData(&IM)

	var tests = []struct {
		name       string
		id         string
		data       models.UpdateArticleModel
		wantError  error
		wantResult models.PackedArticleModel
	}{
		{
			name: "success",
			data: models.UpdateArticleModel{
				ID: "b6c46a35-d78b-4f1d-80fe-8d617f83ec6c",
				Content: models.Content{
					Title: "a",
					Body:  "b",
				},
			},
			wantError: nil,
			wantResult: models.PackedArticleModel{
				Content: models.Content{
					Title: "a",
					Body:  "b",
				},
			},
		},
		{
			name: "fail",
			data: models.UpdateArticleModel{
				ID:      "e87610b5-75ee-4992-99f4-fcff0ca6b4a2",
				Content: models.Content{},
			},
			wantError:  errors.New("article not found"),
			wantResult: models.PackedArticleModel{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IM.UpdateArticle(tt.data)

			if tt.wantError == nil {
				if err != nil {
					t.Errorf("IM.UpdateArticle() got error: %v", err)
				}

				article, err := IM.GetArticleByID(tt.data.ID)
				if err != nil {
					t.Errorf("IM.UpdateArticle() got unexpected error: %v", err)
				}

				if tt.wantResult.Content != article.Content {
					t.Errorf("IM.UpdateArticle() expected: %+v but got: %+v", tt.wantResult.Content, article.Content)
				}
			} else {
				if tt.wantError.Error() != err.Error() {
					t.Errorf("IM.UpdateArticle() expected error: %v, but got error: %v", tt.wantError, err)
				}
			}
		})
	}

	t.Log("Test has been finished")
}
