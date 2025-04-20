package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID 生成一个唯一的ID
func GenerateID(prefix string) string {
	b := make([]byte, 16)
	rand.Read(b)
	return prefix + "_" + hex.EncodeToString(b)
}
