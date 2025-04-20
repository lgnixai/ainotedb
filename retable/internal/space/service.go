package space

import (
	"errors"
	"gorm.io/gorm"
)

type SpaceService struct {
	db *gorm.DB
}

func NewSpaceService(db *gorm.DB) *SpaceService {
	return &SpaceService{db: db}
}

func (s *SpaceService) CreateSpace(space *Space) error {
	return s.db.Create(space).Error
}

func (s *SpaceService) GetSpace(id string) (*Space, error) {
	var space Space
	if err := s.db.First(&space, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &space, nil
}

func (s *SpaceService) AddMember(member *SpaceMember) error {
	return s.db.Create(member).Error
}