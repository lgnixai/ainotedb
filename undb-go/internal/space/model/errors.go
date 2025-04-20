package model

import "errors"

var (
	// ErrEmptySpaceName 空间名称为空
	ErrEmptySpaceName = errors.New("space name is empty")

	// ErrEmptyOwnerID 所有者ID为空
	ErrEmptyOwnerID = errors.New("owner ID cannot be empty")

	// ErrInvalidVisibility is returned when space visibility is invalid
	ErrInvalidVisibility = errors.New("invalid space visibility")

	// ErrEmptySpaceID is returned when space ID is empty
	ErrEmptySpaceID = errors.New("space ID cannot be empty")

	// ErrEmptyUserID is returned when user ID is empty
	ErrEmptyUserID = errors.New("user ID cannot be empty")

	// ErrEmptyRole is returned when role is empty
	ErrEmptyRole = errors.New("role cannot be empty")

	// ErrMemberNotFound is returned when member is not found
	ErrMemberNotFound = errors.New("member not found")

	// ErrInvalidRole is returned when role is invalid
	ErrInvalidRole = errors.New("invalid role")

	// ErrSpaceNotFound 空间不存在
	ErrSpaceNotFound = errors.New("space not found")
)
