package persistence

import (
	"best-practics/common/config"
	"best-practics/common/initialize/log"
	"best-practics/domain/entity"
	"best-practics/domain/interface/repository"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories() (*Repositories) {
	// 初始化数据库连接
	db := InitDB()
	InitTables(db)
	// Dao层方法集合，用于注入给上层
	return &Repositories{
		User: NewUserDao(db),
		db:   db,
	}
}

func InitDB() *gorm.DB {
	m := config.ConfigCenter.Mysql
	if m.Dbname == "" {
		log.Error("未加载MySQL配置")
		os.Exit(0)
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用`change`重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		log.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//注册数据库表
func InitTables(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.BlogType{},
		entity.User{},
	)
	if err != nil {
		log.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	log.Info("register table success")
}


// 程序结束前关闭数据库链接
func (s *Repositories) Close() error {
	db, _ := s.db.DB()
	return db.Close()
}