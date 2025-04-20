package model

// FilterOperator 定义过滤操作符
type FilterOperator string

const (
	OpEqual              FilterOperator = "eq"
	OpNotEqual           FilterOperator = "neq"
	OpGreaterThan        FilterOperator = "gt"
	OpLessThan           FilterOperator = "lt"
	OpGreaterThanOrEqual FilterOperator = "gte"
	OpLessThanOrEqual    FilterOperator = "lte"
	OpContains           FilterOperator = "contains"
	OpNotContains        FilterOperator = "not_contains"
	OpIsEmpty            FilterOperator = "is_empty"
	OpIsNotEmpty         FilterOperator = "is_not_empty"
)

// FilterCondition 表示一个过滤条件
type FilterCondition struct {
	FieldID  string         `json:"fieldId"`
	Operator FilterOperator `json:"operator"`
	Value    interface{}    `json:"value,omitempty"`
}

// FilterGroup 表示一组过滤条件
type FilterGroup struct {
	Type       string            `json:"type"` // and, or
	Conditions []FilterCondition `json:"conditions"`
	Groups     []FilterGroup     `json:"groups,omitempty"`
}

// FormulaFilter 表示公式过滤
type FormulaFilter struct {
	Formula string `json:"formula"` // 例如: {field1} > 100 AND {field2} CONTAINS 'test'
}

// Validate 验证过滤条件
func (fc *FilterCondition) Validate() error {
	if fc.FieldID == "" {
		return ErrEmptyFieldID
	}
	// 验证操作符是否支持
	switch fc.Operator {
	case OpEqual, OpNotEqual, OpGreaterThan, OpLessThan,
		OpGreaterThanOrEqual, OpLessThanOrEqual,
		OpContains, OpNotContains, OpIsEmpty, OpIsNotEmpty:
		return nil
	default:
		return ErrInvalidFilterOperator
	}
}
