package models

import "time"

// Person ...
type Person struct {
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
}

// Content ...
type Content struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// Article ...
type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author    Person     `json:"a" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateArticleModel ...
type CreateArticleModel struct {
	Content        // Promoted fields
	Author  Person `json:"a" binding:"required"`
}

// JSONResponse ...
type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSONErrorResponse ...
type JSONErrorResponse struct {
	Error string `json:"error"`
}
