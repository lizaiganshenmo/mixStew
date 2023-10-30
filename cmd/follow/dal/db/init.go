package db

import (
	config "github.com/lizaiganshenmo/mixStew/cmd/follow/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	MySQLMixStewDB *gorm.DB
)

func Init() {
	initSQL("mysql_mixStew")
}

// init mysql
func initSQL(srvName string) {
	var dsn string
	var err error
	dsn, err = config.GetMysqlDSN(srvName)
	if err != nil {
		panic(err)
	}

	MySQLMixStewDB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true, // 禁用默认事务
		},
	)
	if err != nil {
		panic(err)
	}

	if err = MySQLMixStewDB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

}
