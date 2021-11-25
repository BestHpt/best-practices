/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package mysql

import (
	"best-practics/common"
	"best-practics/common/initialize/log"
	"best-practics/common/initialize/viper"
	"best-practics/domain/entity"
	"database/sql"
	"fmt"
	"testing"
)

var db *sql.DB

// 这里不用小写默认，防止再测试其他单测文件时也跑到这里
func Init() {
	common.Viper = viper.Init("../../conf/logConfig.yaml") // 初始化Viper
	log.InitZap()                                          // 初始化zap日志库
	common.GVA_DB = Init()                                 // gorm连接数据库
	if common.GVA_DB != nil {
		MysqlTables(common.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ = common.GVA_DB.DB()
	}
}

func TestGorm(t *testing.T) {
	Init()
	defer db.Close()
	var blogType entity.BlogType
	err := common.GVA_DB.Where("id = ?", 1).First(&blogType).Error
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("value is:", blogType)
	}
}

type ColumnReq struct {
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataType      string `json:"dataType" gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}


// 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName         string   `json:"structName"`         // Struct名称
	TableName          string   `json:"tableName"`          // 表名
	PackageName        string   `json:"packageName"`        // 文件名称
	HumpPackageName    string   `json:"humpPackageName"`    // go文件名称
	Abbreviation       string   `json:"abbreviation"`       // Struct简称
	Description        string   `json:"description"`        // Struct中文名称
	AutoCreateApiToSql bool     `json:"autoCreateApiToSql"` // 是否自动创建api
	AutoMoveFile       bool     `json:"autoMoveFile"`       // 是否自动移动文件
	Fields             []*Field `json:"fields"`
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataType        string `json:"dataType"`        // 数据库字段类型
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	DictType        string `json:"dictType"`        // 字典
}

func TestAutoCode(t *testing.T) {
	Init()
	defer db.Close()
	tableName := "blog"
	dbName := "blog"
	var Columns []ColumnReq
	err := common.GVA_DB.Raw("SELECT COLUMN_NAME column_name,DATA_TYPE data_type,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c." +
		"CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT column_comment " +
		"FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", tableName, dbName).Scan(&Columns).Error

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("value is:", Columns)
	}
}

