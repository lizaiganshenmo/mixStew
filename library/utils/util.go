package utils

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
)

func EncryptPassword(password string) string {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		klog.Warnf("EncryptPassword fail. password: %s", password)
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
