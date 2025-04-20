package model

import "errors"

var (
	// ErrEmptyTableID 表ID为空
	ErrEmptyTableID = errors.New("table ID cannot be empty")
	// ErrEmptyFields 记录字段为空
	ErrEmptyFields = errors.New("record fields cannot be empty")
	// ErrRecordNotFound 记录不存在
	ErrRecordNotFound = errors.New("record not found")
)
