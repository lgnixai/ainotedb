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

func (s *SpaceService) UpdateSpace(space *Space) error {
	return s.db.Save(space).Error
}

func (s *SpaceService) DeleteSpace(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Delete space members
		if err := tx.Where("space_id = ?", id).Delete(&SpaceMember{}).Error; err != nil {
			return err
		}
		// Delete the space
		return tx.Delete(&Space{}, "id = ?", id).Error
	})
}

func (s *SpaceService) GetMembers(spaceID string) ([]SpaceMember, error) {
	var members []SpaceMember
	err := s.db.Where("space_id = ?", spaceID).Find(&members).Error
	return members, err
}