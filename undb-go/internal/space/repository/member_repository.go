package repository

import (
	"context"

	"github.com/undb/undb-go/internal/space/model"
	"gorm.io/gorm"
)

// MemberRepository defines the interface for space member data access
type MemberRepository interface {
	// Create creates a new space member
	Create(ctx context.Context, member *model.SpaceMember) error

	// GetByID retrieves a member by its ID
	GetByID(ctx context.Context, id string) (*model.SpaceMember, error)

	// GetBySpaceID retrieves all members of a space
	GetBySpaceID(ctx context.Context, spaceID string) ([]*model.SpaceMember, error)

	// GetByUserID retrieves all spaces a user is a member of
	GetByUserID(ctx context.Context, userID string) ([]*model.SpaceMember, error)

	// Update updates a member's role
	Update(ctx context.Context, member *model.SpaceMember) error

	// Delete removes a member from a space
	Delete(ctx context.Context, spaceID, userID string) error

	// UpdateRole updates a member's role
	UpdateRole(ctx context.Context, spaceID, userID string, role model.MemberRole) error
}

// memberRepository implements the MemberRepository interface
type memberRepository struct {
	db *gorm.DB
}

// NewMemberRepository creates a new member repository instance
func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) Create(ctx context.Context, member *model.SpaceMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

func (r *memberRepository) GetByID(ctx context.Context, id string) (*model.SpaceMember, error) {
	var member model.SpaceMember
	if err := r.db.WithContext(ctx).First(&member, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) GetBySpaceID(ctx context.Context, spaceID string) ([]*model.SpaceMember, error) {
	var members []*model.SpaceMember
	if err := r.db.WithContext(ctx).Where("space_id = ?", spaceID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) GetByUserID(ctx context.Context, userID string) ([]*model.SpaceMember, error) {
	var members []*model.SpaceMember
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) Update(ctx context.Context, member *model.SpaceMember) error {
	return r.db.WithContext(ctx).Save(member).Error
}

func (r *memberRepository) Delete(ctx context.Context, spaceID, userID string) error {
	return r.db.WithContext(ctx).Delete(&model.SpaceMember{}, "space_id = ? AND user_id = ?", spaceID, userID).Error
}

func (r *memberRepository) UpdateRole(ctx context.Context, spaceID, userID string, role model.MemberRole) error {
	return r.db.WithContext(ctx).Model(&model.SpaceMember{}).
		Where("space_id = ? AND user_id = ?", spaceID, userID).
		Update("role", role).Error
}
