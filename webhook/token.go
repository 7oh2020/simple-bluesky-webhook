package webhook

import (
	"crypto/rand"
	"encoding/base64"
)

// URLセーフなランダム文字列を返します。
func GenerateToken(size int) (string, error) {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)

	return token, nil
}
