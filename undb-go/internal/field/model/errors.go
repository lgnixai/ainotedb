package model

import "errors"

var (
	// ErrEmptyFieldName 字段名称为空
	ErrEmptyFieldName = errors.New("field name cannot be empty")
	// ErrEmptyFieldType 字段类型为空
	ErrEmptyFieldType = errors.New("field type cannot be empty")
	// ErrInvalidFieldType 字段类型无效
	ErrInvalidFieldType = errors.New("invalid field type")
	// ErrFieldNotFound 字段不存在
	ErrFieldNotFound = errors.New("field not found")
)
