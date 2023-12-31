package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func DataBase(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,   // string类型字段默认长度
		DisableDatetimePrecision:  true,  // 禁止datetime精度 mysql 5.6之前是不支持
		DontSupportRenameIndex:    true,  //重命名索引 要先删除再重建 mysql 5.7之前是不支持的
		DontSupportRenameColumn:   true,  // 用change重命名列 mysql8之前不支持
		SkipInitializeWithVersion: false, // 不 跳过 Gorm 初始化检查
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表不加s
		},
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20) //连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	sqlDB.SetMaxIdleConns(100) // 设置连接池
	_db = db

	//	主从配置
	_ = _db.Use(dbresolver.
		Register(
			dbresolver.Config{
				Sources:  []gorm.Dialector{mysql.Open(connWrite)},                       // 写操作
				Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connWrite)}, //读操作
				Policy:   dbresolver.RandomPolicy{},
			}))
	migration()
}

func NewDbClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
