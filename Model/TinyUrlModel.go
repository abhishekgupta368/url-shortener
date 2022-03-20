package model

import "gorm.io/gorm"

type UrlModel struct {
	Url string `json:"url"`
}

type TinyUrlData struct {
	gorm.Model
	Hash string
	Url  string
}
