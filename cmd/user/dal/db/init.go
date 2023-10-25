package db

import (
	"fmt"

	config "github.com/lizaiganshenmo/mixStew/cmd/user/configs"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	MySQLMixStewDB *gorm.DB
	SF             *utils.Snowflake
)

func Init() {
	initSQL("mysql_mixStew")
	initSF()
}

// init mysql
func initSQL(srvName string) {
	var dsn string
	var err error
	dsn, err = config.GetMysqlDSN(srvName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sql dsn is: %s\n", dsn)

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

// init snowflake
func initSF() {
	var err error
	if SF, err = utils.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID); err != nil {
		panic(err)
	}
}