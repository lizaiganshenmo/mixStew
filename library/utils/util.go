package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptPassword(password string) string {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
