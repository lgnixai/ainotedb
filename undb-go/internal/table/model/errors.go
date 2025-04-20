package model

import "errors"

var (
	// ErrEmptyTableName 表格名称为空
	ErrEmptyTableName = errors.New("table name cannot be empty")

	// ErrEmptyTableID 表格ID为空
	ErrEmptyTableID = errors.New("table ID cannot be empty")

	// ErrEmptySpaceID 空间ID为空
	ErrEmptySpaceID = errors.New("space ID cannot be empty")

	// ErrEmptyFieldName 字段名称为空
	ErrEmptyFieldName = errors.New("field name cannot be empty")

	// ErrEmptyFieldType 字段类型为空
	ErrEmptyFieldType = errors.New("field type cannot be empty")

	// ErrEmptyFields 记录字段为空
	ErrEmptyFields = errors.New("record fields cannot be empty")

	// ErrInvalidFieldType 无效的字段类型
	ErrInvalidFieldType = errors.New("invalid field type")

	// ErrTableNotFound 表格不存在
	ErrTableNotFound = errors.New("table not found")

	// ErrFieldNotFound 字段不存在
	ErrFieldNotFound = errors.New("field not found")

	// ErrRecordNotFound 记录不存在
	ErrRecordNotFound = errors.New("record not found")
)
