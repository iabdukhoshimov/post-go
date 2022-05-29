package models

import "time"

type Person struct {
	ID        string     `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PersonCreateModel struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type PersonUpdateModel struct {
	ID        string `json:"-"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
