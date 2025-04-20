package repository

import (
	"context"
	"database/sql"

	"github.com/undb/undb-go/internal/space/model"
	"gorm.io/gorm"
)

type MemberRepository interface {
	Create(ctx context.Context, member *model.SpaceMember) error
	Update(ctx context.Context, member *model.SpaceMember) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.SpaceMember, error)
	FindBySpaceID(ctx context.Context, spaceID string) ([]*model.SpaceMember, error)
	FindByUserID(ctx context.Context, userID string) ([]*model.SpaceMember, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) Create(ctx context.Context, member *model.SpaceMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

func (r *memberRepository) Update(ctx context.Context, member *model.SpaceMember) error {
	return r.db.WithContext(ctx).Save(member).Error
}

func (r *memberRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.SpaceMember{}, "id = ?", id).Error
}

func (r *memberRepository) FindByID(ctx context.Context, id string) (*model.SpaceMember, error) {
	var member model.SpaceMember
	if err := r.db.WithContext(ctx).First(&member, "id = ?", id).Error; err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) FindBySpaceID(ctx context.Context, spaceID string) ([]*model.SpaceMember, error) {
	var members []*model.SpaceMember
	if err := r.db.WithContext(ctx).Find(&members, "space_id = ?", spaceID).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) FindByUserID(ctx context.Context, userID string) ([]*model.SpaceMember, error) {
	var members []*model.SpaceMember
	if err := r.db.WithContext(ctx).Find(&members, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return members, nil
}