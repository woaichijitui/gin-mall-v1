package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var DB *gorm.DB

func Database(connRead, connWrite string) {
	//1日志设置
	var ormlogger logger.Interface
	if gin.Mode() == "debug" {
		ormlogger = logger.Default.LogMode(logger.Info)
	} else {
		ormlogger = logger.Default
	}

	//2、开启数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      connRead,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
		DontSupportRenameColumn:  true,

		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger:                                   ormlogger,
		DisableForeignKeyConstraintWhenMigrating: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("database open err :", err)
	}

	//3、数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("sql pool err:", err)
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("success to link mysql")
	DB = db

	//4、主从复制
	_ = DB.Use(dbresolver.Register(dbresolver.Config{
		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
		Sources:  []gorm.Dialector{mysql.Open(connRead)},                         // 写操作
		Replicas: []gorm.Dialector{mysql.Open(connWrite), mysql.Open(connWrite)}, // 读操作
		Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas 负载均衡策略))
	}))

	//Migration()
}
func NewDBClient(context context.Context) *gorm.DB {

	db := DB
	return db.WithContext(context)
}
