package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

func GenerateUserID(pubKey []byte) string {
	hash := sha256.Sum256(pubKey)
	short := fmt.Sprintf("%x", hash[:8]) // 16 hex chars
	b58 := base58.Encode([]byte(short))
	checksum := "7X" // Simple MVP
	rawID := fmt.Sprintf("%s-%s", b58, checksum)
	return fmt.Sprintf("ora:%s-%s-%s-%s", rawID[0:4], rawID[4:8], rawID[8:12], rawID[len(rawID)-2:])
}
