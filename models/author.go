package models

import "time"

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Article struct {
	ID        int        `json:"id"`
	Content              //promoted fields
	AuthorID  int        `json:"author_id"`
	CreatedAt *time.Time `json:"created_at:"`
	UpdateAt  *time.Time `json:"updated_at"`
}

type ArticleListItem struct {
	ID        int        `json:"id"`
	Content              //promoted fields
	Author    Person     `json:"author"`
	CreatedAt *time.Time `json:"created_at:"`
	UpdateAt  *time.Time `json:"updated_at"`
}

type ArticleCreateModel struct {
	Content      //promoted fields
	AuthorID int `json:"author_id"`
}

type ArticleUpdateModel struct {
	ID       int `json:"-"`
	Content      //promoted fields
	AuthorID int `json:"author_id"`
}
