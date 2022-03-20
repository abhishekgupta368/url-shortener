package repository

import (
	model "TinyUrl/Model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&model.TinyUrlData{})
	return &Repository{
		Db: db,
	}
}

func (r *Repository) Save(urlModel model.UrlModel, hash string) {
	var urlData model.TinyUrlData
	r.Db.Where("hash = (?) ", hash).Find(&urlData)
	if urlData.Url == "" {
		r.Db.Create(&model.TinyUrlData{
			Hash: hash,
			Url:  urlModel.Url,
		})
		log.Println("Data is created")
	} else {
		id := uuid.New()
		r.Db.Create(&model.TinyUrlData{
			Hash: hash,
			Url:  id.String()[:8],
		})
		log.Println("Hash already exists, creating other hash")
	}
}

func (r *Repository) Get(hash string) model.TinyUrlData {
	var urlData model.TinyUrlData
	r.Db.Where("hash = (?) ", hash).Find(&urlData)
	return urlData
}
