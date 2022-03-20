package service

import (
	model "TinyUrl/Model"
	repository "TinyUrl/Repository"
	utility "TinyUrl/Utility"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

var (
	UrlContent = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	UrlAddress = "localhost:8080/"
)

func (s *Service) GenerateTinyUrl(urlModel model.UrlModel) model.UrlModel {
	hash := utility.ComputeHash(urlModel.Url)
	len := len(UrlContent)
	hashUrl := ""
	for hash > 0 {
		idx := hash % int64(len)
		hash = hash / int64(len)
		hashUrl += string(UrlContent[idx])
	}
	computedUrl := UrlAddress + hashUrl
	s.repo.Save(urlModel, hashUrl)
	return model.UrlModel{
		Url: computedUrl,
	}
}

func (s *Service) RedirectTinyUrl(hash string) string {
	tinyUrl := s.repo.Get(hash)
	return tinyUrl.Url
}
