package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
	arrClass := make([]model.Class, 0)
	data := s.db.Find(&arrClass)
	if data.Error != nil {
		return nil, data.Error
	}
	return arrClass, nil // TODO: replace this
}
