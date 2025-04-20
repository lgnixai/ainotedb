
package space

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpaceService struct {
	db *gorm.DB
}

func NewSpaceService(db *gorm.DB) *SpaceService {
	return &SpaceService{db: db}
}

func (s *SpaceService) CreateSpace(ctx context.Context, space *Space) error {
	if space.ID == "" {
		space.ID = uuid.New().String()
	}
	
	return s.db.WithContext(ctx).Create(space).Error
}

func (s *SpaceService) UpdateSpace(ctx context.Context, space *Space) error {
	if space.ID == "" {
		return errors.New("space ID is required")
	}
	
	return s.db.WithContext(ctx).Save(space).Error
}

func (s *SpaceService) DeleteSpace(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete space members first
		if err := tx.WithContext(ctx).Where("space_id = ?", id).Delete(&SpaceMember{}).Error; err != nil {
			return err
		}
		
		// Delete the space
		return tx.WithContext(ctx).Delete(&Space{}, "id = ?", id).Error
	})
}

func (s *SpaceService) GetSpace(ctx context.Context, id string) (*Space, error) {
	var space Space
	if err := s.db.WithContext(ctx).First(&space, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &space, nil
}

func (s *SpaceService) ListSpaces(ctx context.Context, userID string) ([]Space, error) {
	var spaces []Space
	err := s.db.WithContext(ctx).
		Joins("JOIN space_members ON spaces.id = space_members.space_id").
		Where("space_members.user_id = ?", userID).
		Find(&spaces).Error
	return spaces, err
}

func (s *SpaceService) AddMember(ctx context.Context, member *SpaceMember) error {
	if member.ID == "" {
		member.ID = uuid.New().String()
	}
	return s.db.WithContext(ctx).Create(member).Error
}

func (s *SpaceService) RemoveMember(ctx context.Context, spaceID, userID string) error {
	return s.db.WithContext(ctx).
		Where("space_id = ? AND user_id = ?", spaceID, userID).
		Delete(&SpaceMember{}).Error
}

func (s *SpaceService) UpdateMemberRole(ctx context.Context, spaceID, userID string, role Role) error {
	return s.db.WithContext(ctx).
		Model(&SpaceMember{}).
		Where("space_id = ? AND user_id = ?", spaceID, userID).
		Update("role", role).Error
}

func (s *SpaceService) GetMembers(ctx context.Context, spaceID string) ([]SpaceMember, error) {
	var members []SpaceMember
	err := s.db.WithContext(ctx).Where("space_id = ?", spaceID).Find(&members).Error
	return members, err
}
