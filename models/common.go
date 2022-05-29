package models

type DefaultError struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    interface{} `json:"data"`
}

type Query struct {
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"Limit" default:"0"`
	Search string `json:"search"`
}
