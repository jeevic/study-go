package models

import _ "github.com/jinzhu/gorm/dialects/mysql"

type Author struct {
	AuthorId      int
	Day           int
	PlatfromId    int
	IncrContent   int
	IncrFollowers int
	IncrPlay      int
	IncrComment   string
	CreatedAt     string
}

func (Author) TableName() string {
	return "analysis_author"
}
