package model

type Note struct {
	Id      int    `gorm:"int;primary_key"`
	Content string `gorm:"not null" json:"content,omitempty"`
}
