package db

import (
	"github.com/atpuxiner/gin-layout/app/datatype/model"
	"github.com/atpuxiner/gin-layout/app/initializer/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func InitDB(driver string) {
	var (
		err       error
		dialector gorm.Dialector
	)
	dsn := conf.Conf.Db.Dsn
	switch driver {
	case "postgres":
		dialector = postgres.Open(dsn)
	case "mysql":
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		panic("暂不支持该数据库驱动: " + driver)
	}
	DB, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	// 配置相关
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(conf.Conf.Db.MaxIdleCount)
	sqlDB.SetMaxOpenConns(conf.Conf.Db.MaxOpenCount)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 迁移数据库
	autoMigrate(DB)
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
}
