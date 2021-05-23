package data

import (
	"testing"
	"github.com/tietang/props/v3/yam"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"context"
	ent2 "github.com/blackbinbinbinbin/Bingo-gin/pkg/ent"
	"log"
)

func getDb() *sql.DB {
	props := yam.NewYamlConfigSource("/Go/Bingo-gin/configs/config.yaml")
	env := props.GetDefault("app_env_mode", "develop")
	ip := props.GetDefault(fmt.Sprintf("mysql.%s.ip", env), "127.0.0.1")
	port := props.GetDefault(fmt.Sprintf("mysql.%s.port", env), "3306")
	user := props.GetDefault(fmt.Sprintf("mysql.%s.user", env), "")
	pwd := props.GetDefault(fmt.Sprintf("mysql.%s.pwd", env), "")
	database := props.GetDefault("mysql.database", "")
	dbConfigStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, ip, port, database)

	dbconn, err := sql.Open("mysql", dbConfigStr)
	if err != nil {
		panic("connect db error")
	}

	return dbconn
}

func TestCreateUser(t *testing.T) {
	db := getDb()
	client, _ := ent2.Open(db)

	u, err := CreateUser(context.Background(), client, 33, "a8m")
	if err != nil {
		panic(err)
	}

	log.Println(u)
}


func TestUpdateUserById(t *testing.T) {
	db := getDb()
	client, _ := ent2.Open(db)

	u, err := UpdateUserById(context.Background(), client, 1, 28, "a7m")
	if err != nil {
		panic(err)
	}
	log.Println(u)
}


func TestQueryUserById(t *testing.T) {
	db := getDb()
	client, _ := ent2.Open(db)

	u, err := QueryUserById(context.Background(), client, 1)
	if err != nil {
		panic(err)
	}

	log.Println(u)
}