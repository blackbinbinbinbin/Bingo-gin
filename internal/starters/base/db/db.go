package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	"fmt"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"time"
)

var dbbase *sql.DB

func DbxDataBase() *sql.DB {
	starters.Check(dbbase)
	return dbbase
}

type DbxDataBaseStarter struct {
	starters.BaseStarter
}

func (d *DbxDataBaseStarter) Setup(ctx starters.StarterContext) {
	// 数据库配置
	props := ctx.Props()
	env := props.GetDefault("app_env_mode", "develop")
	ip := props.GetDefault(fmt.Sprintf("mysql.%s.ip", env), "127.0.0.1")
	port := props.GetDefault(fmt.Sprintf("mysql.%s.port", env), "3306")
	user := props.GetDefault(fmt.Sprintf("mysql.%s.user", env), "")
	pwd := props.GetDefault(fmt.Sprintf("mysql.%s.pwd", env), "")
	database := props.GetDefault("mysql.database", "")
	dbConfigStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, ip, port, database)

	l := log.DefaultLogger
	log.Info(l).Log("msg", "连接gorm，配置"+dbConfigStr)

	dbconn, err := sql.Open("mysql", dbConfigStr)
	if err != nil {
		log.Error(l).Log("msg", "数据库初始化错误，连接出错！！")
		panic("数据库初始化错误，连接出错！！")
		return
	}

	dbconn.SetMaxOpenConns(2000)
	dbconn.SetMaxIdleConns(1000)
	dbconn.SetConnMaxLifetime(time.Hour)
	if err := dbconn.Ping(); err != nil {
		log.Error(l).Log("msg", "数据库初始化错误，连接出错！！")
		panic("数据库初始化错误，连接出错！！")
		return
	}

	dbbase = dbconn
}