package model

import "github.com/undb/undb-go/pkg/utils"

// GenerateID 生成唯一ID
func GenerateID() string {
	return utils.GenerateID("tbl")
}
