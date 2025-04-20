package model

import "errors"

var (
	// ErrEmptyViewName is returned when view name is empty
	ErrEmptyViewName = errors.New("view name is required")
	// ErrEmptyTableID is returned when table ID is empty
	ErrEmptyTableID = errors.New("table ID is required")
	// ErrEmptyViewType is returned when view type is empty
	ErrEmptyViewType = errors.New("view type is required")
	// ErrViewNotFound is returned when view is not found
	ErrViewNotFound = errors.New("view not found")
	// ErrInvalidViewType is returned when view type is invalid
	ErrInvalidViewType = errors.New("invalid view type")
)
