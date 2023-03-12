package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type TagService interface {
	GetAllTags() ([]models.Tag, error)
	CreateTag(models.Tag) (models.Tag, error)
	UpdateTag(models.Tag) (models.Tag, error)
	DeleteTag(string) error
}

type tagService struct {
	tagRepo repositories.TagRepo
}

func NewTagService(tagRepo repositories.TagRepo) TagService {
	return &tagService{tagRepo: tagRepo}
}

func (s *tagService) GetAllTags() ([]models.Tag, error) {
	return s.tagRepo.GetAllTags()
}

func (s *tagService) CreateTag(newTag models.Tag) (models.Tag, error) {
	return s.tagRepo.CreateTag(newTag)
}

func (s *tagService) UpdateTag(newTag models.Tag) (models.Tag, error) {
	return s.tagRepo.UpdateTag(newTag)
}

func (s *tagService) DeleteTag(id string) error {
	return s.tagRepo.DeleteTag(id)
}
