package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	userName = "root"
	password = "12345678"
	path     = "localhost:3306"
	dbName   = "education"
	// Db 全局db
	Db *gorm.DB
)

func Init() {
	initMysql()
}

func initMysql() {

	dsn := userName + ":" + password + "@tcp(" + path + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	var err error
	if Db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		fmt.Println("MySQL启动异常:" + err.Error())
		panic(err)
	}
}
