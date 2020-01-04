package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"reflect"
	"Bingo-gin/common"
	"time"
)

type MysqlDb struct {
	// 数据库连接对象
	dbconn *gorm.DB

	// 库名
	Dababase string

	// 表名
	Table string

	// 调试模式
	DebugMode bool

	// 实体
	Entity struct{}
}

var DBCon *gorm.DB

func (DB *MysqlDb) Init() interface{}{
	var err error
	if DBCon != nil && DBCon.DB().Ping() != nil {
		return DB
	}

	// 数据相关配置
	cfg := common.NewConfig()
	appMode := cfg.GetValue("", "APP_ENV_MODE")
	ip := cfg.GetSectionValue(appMode, "MYSQL_SERVER_IP", "127.0.0.1")
	port := cfg.GetSectionValue(appMode, "MYSQL_SERVER_PORT", "3306")
	user := cfg.GetSectionValue(appMode, "MYSQL_SERVER_USER", "")
	pwd := cfg.GetSectionValue(appMode, "MYSQL_SERVER_PWD", "")
	dbConfigStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, ip, port, DB.Dababase)

	common.Info("连接gorm，配置"+fmt.Sprintf("%s:***@tcp(%s:%s)/%s?charset=utf8", user, ip, port, DB.Dababase))
	DB.dbconn, err = gorm.Open("mysql", dbConfigStr)
	DBCon = DB.dbconn
	// 设置空闲连接池中的最大连接数
	DB.dbconn.DB().SetMaxIdleConns(10)
	// 设置到数据库的最大打开连接数
	DB.dbconn.DB().SetMaxOpenConns(100)
	// 设置可以重用连接的最长时间
	DB.dbconn.DB().SetConnMaxLifetime(time.Hour)
	if DB.DebugMode {
		DB.dbconn.LogMode(DB.DebugMode)
	}

	if err != nil {
		fmt.Println("failed to connect database:", err)
		return err
	}

	return DB
}


// 根据条件查找数据
//	example:
//  	var rowsEntity []entity.TableEntity
//		DB.Select(&rowsEntity, where, limit, 0, "*", "id")
func (DB *MysqlDb) Select (
	entity interface{},
	where map[string]interface{},
	limit int, offset int, fields string, orderBy string) interface{} {

	// 处理 where 参数
	objDb := DB.dbconn.Table(DB.Table)
	for fieldName, val := range where {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			objDb = objDb.Where(fieldName+" IN (?)", val)
		} else {
			objDb = objDb.Where(fieldName+" = ?", val)
		}
	}

	rows := objDb.Limit(limit).Offset(offset).Order(orderBy)
	if result := rows.Find(entity); result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}

	return entity
}


// 执行查询sql语句
//	example:
//  	var rowEntity entity.TableEntity
//	    DB.QuerySql(&rowEntity,"select * from `table` where id in (?)", 2)
func (DB *MysqlDb) QuerySql(
	entity interface{},
	sql string, values ...interface{}) interface{} {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Raw(sql, values).Find(entity)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
	return entity
}


// 执行原生的sql语句
//	example:
//	    DB.ExecSql("DROP TABLE table", nil)
//	    DB.ExecSql("UPDATE orders SET shipped_at=? WHERE id IN (?)", time.Now(), []int64{11,22,33})
func (DB *MysqlDb) ExecSql(sql string, values ...interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Exec(sql, values)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 插入数据
//	example:
//		var table = entity.TableEntity{fieldName: value}
//		DB.Insert(&table)
func (DB *MysqlDb) Insert(entity interface{}) interface{} {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Create(entity)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
	return entity
}


// 根据实体来修改某一行数据
//	example:
//		var updateData = map[string]interface{}{"fieldName": "value"}
//		DB.UpdateOneRow(&tableEntity, updateData)
func (DB *MysqlDb) UpdateOneRow(entity interface{}, updateData map[string]interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Model(entity).Updates(updateData)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 根据条件查询修改多行数据
//	example:
//		var where = map[string]interface{}{"id": []int{1, 2}}
//		var updateData = map[string]interface{}{"fieldName": "value"}
//		DB.UpdateBatchRows(where, updateData)
func (DB *MysqlDb) UpdateBatchRows(where map[string]interface{}, updateData map[string]interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	for fieldName, val := range where {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			objDb = objDb.Where(fieldName+" IN (?)", val)
		} else {
			objDb = objDb.Where(fieldName+" = ?", val)
		}
	}

	result := objDb.Updates(updateData)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 根据实体entity删除某行数据
//	example:
//		DB.DeleteOneRow(&tableEntity)
func (DB *MysqlDb) DeleteOneRow(entity interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Delete(&entity)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 根据搜索条件删除多行
//	example:
//		var where = map[string]interface{}{"fieldName": "value"}
//		DB.DeleteBatchRows(where)
func (DB *MysqlDb) DeleteBatchRows(where map[string]interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	for fieldName, val := range where {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			objDb = objDb.Where(fieldName+" IN (?)", val)
		} else {
			objDb = objDb.Where(fieldName+" = ?", val)
		}
	}
	result := objDb.Delete(nil)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}

