package models

type Post struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
