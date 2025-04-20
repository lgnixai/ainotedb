package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/undb/undb-go/internal/space/model"
	"github.com/undb/undb-go/internal/space/repository"
	"github.com/undb/undb-go/pkg/utils"
)

// SpaceService 定义空间服务接口
type SpaceService interface {
	// Create 创建新空间
	Create(ctx context.Context, space *model.Space) error

	// GetByID 根据ID获取空间
	GetByID(ctx context.Context, id string) (*model.Space, error)

	// GetByOwnerID 获取用户的所有空间
	GetByOwnerID(ctx context.Context, ownerID string) ([]*model.Space, error)

	// Update 更新空间
	Update(ctx context.Context, space *model.Space) error

	// Delete 删除空间
	Delete(ctx context.Context, id string) error

	// Member management
	AddMember(ctx context.Context, spaceID, userID string, role model.MemberRole) error
	RemoveMember(ctx context.Context, spaceID, userID string) error
	UpdateMemberRole(ctx context.Context, spaceID, userID string, role model.MemberRole) error
	GetSpaceMembers(ctx context.Context, spaceID string) ([]*model.SpaceMember, error)
}

// spaceService 实现空间服务接口
type spaceService struct {
	spaceRepo  repository.SpaceRepository
	memberRepo repository.MemberRepository
	db         *gorm.DB
}

// NewSpaceService 创建新的空间服务实例
func NewSpaceService(spaceRepo repository.SpaceRepository, memberRepo repository.MemberRepository, db *gorm.DB) SpaceService {
	return &spaceService{
		spaceRepo:  spaceRepo,
		memberRepo: memberRepo,
		db:         db,
	}
}

func (s *spaceService) Create(ctx context.Context, space *model.Space) error {
	if err := space.Validate(); err != nil {
		return err
	}

	// 创建空间
	if err := s.spaceRepo.Create(ctx, space); err != nil {
		return err
	}

	// 自动添加创建者为所有者
	member := &model.SpaceMember{
		ID:      utils.GenerateID("mem"),
		SpaceID: space.ID,
		UserID:  space.OwnerID,
		Role:    model.RoleOwner,
	}
	return s.memberRepo.Create(ctx, member)
}

func (s *spaceService) GetByID(ctx context.Context, id string) (*model.Space, error) {
	return s.spaceRepo.GetByID(ctx, id)
}

func (s *spaceService) GetByOwnerID(ctx context.Context, ownerID string) ([]*model.Space, error) {
	return s.spaceRepo.GetByOwnerID(ctx, ownerID)
}

func (s *spaceService) Update(ctx context.Context, space *model.Space) error {
	if err := space.Validate(); err != nil {
		return err
	}
	return s.spaceRepo.Update(ctx, space)
}

func (s *spaceService) Delete(ctx context.Context, id string) error {
	// 先删除所有成员
	if err := s.db.WithContext(ctx).Delete(&model.SpaceMember{}, "space_id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete space members: %w", err)
	}
	// 再删除空间
	return s.spaceRepo.Delete(ctx, id)
}

func (s *spaceService) AddMember(ctx context.Context, spaceID, userID string, role model.MemberRole) error {
	member := &model.SpaceMember{
		ID:      utils.GenerateID("mem"),
		SpaceID: spaceID,
		UserID:  userID,
		Role:    role,
	}

	if err := member.Validate(); err != nil {
		return err
	}

	return s.memberRepo.Create(ctx, member)
}

func (s *spaceService) RemoveMember(ctx context.Context, spaceID, userID string) error {
	return s.memberRepo.Delete(ctx, spaceID)
}

func (s *spaceService) UpdateMemberRole(ctx context.Context, spaceID, userID string, role model.MemberRole) error {
	// 先查找成员
members, err := s.memberRepo.FindBySpaceID(ctx, spaceID)
if err != nil {
    return err
}
var memberToUpdate *model.SpaceMember
for _, m := range members {
    if m.UserID == userID {
        memberToUpdate = m
        break
    }
}
if memberToUpdate == nil {
    return fmt.Errorf("member not found")
}
memberToUpdate.Role = role
return s.memberRepo.Update(ctx, memberToUpdate)
}

func (s *spaceService) GetSpaceMembers(ctx context.Context, spaceID string) ([]*model.SpaceMember, error) {
	return s.memberRepo.FindBySpaceID(ctx, spaceID)
}
