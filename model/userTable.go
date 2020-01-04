package model

import (
	"sync"
	"Bingo-gin/entity"
)

type UserTable struct {
	MysqlDb
}

var instanceUserTable *UserTable
var once sync.Once

/**
 * 单例模式
 */
func UserTableInstance() *UserTable{
	once.Do(func(){
		instanceUserTable = &UserTable{}
		instanceUserTable.New()
	})
	return instanceUserTable
}

func (DB *UserTable) New() *UserTable {
	DB.Dababase = "bingo"
	DB.Table = "user"
	DB.DebugMode = true
	if DBCon == nil || DBCon.DB().Ping() == nil {
		DB.Init()
	}
	return DB
}

func (DB *UserTable) GetUsersById(id []int) []entity.UserEntity{
	where := make(map[string]interface{})
	where["id"] = id
	limit := len(id)
	var users []entity.UserEntity
	DB.Select(&users, where, limit, 0, "*", "id")

	return users
}