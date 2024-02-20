package dal

import (
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/cache"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
