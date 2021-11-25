/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package mysql

import (
	"best-practics/common/config"
	"best-practics/common/initialize/log"
	"best-practics/domain/entity"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Init() *gorm.DB {
	switch config.ConfigCenter.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.BlogType{},
	)
	if err != nil {
		log.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	log.Info("register table success")
}

//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := config.ConfigCenter.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		//log.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	gormConfigInstance := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch config.ConfigCenter.Mysql.LogMode {
	case "silent", "Silent":
		gormConfigInstance.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		gormConfigInstance.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		gormConfigInstance.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		gormConfigInstance.Logger = Default.LogMode(logger.Info)
	default:
		gormConfigInstance.Logger = Default.LogMode(logger.Info)
	}
	return gormConfigInstance
}
