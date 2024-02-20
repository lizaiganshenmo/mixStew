package dal

import (
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/mq"
	sensitiveword "github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/sensitive_word"
)

func Init(path string) {
	// 敏感词初始化
	sensitiveword.Init("/Users/saiyajin/Desktop/now/mixStew/cmd/interaction/configs/sensitive_word/")
	db.Init()
	mq.Init()
}
