package main

import (
	"github.com/tietang/props/v3/yam"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	_ "github.com/blackbinbinbinbin/Bingo-gin/cmd/my-app/init"
	"github.com/blackbinbinbinbin/Bingo-gin/pkg/ent"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/data"
	"context"
	"fmt"
)

func main() {
	// 加载配置文件
	conf := yam.NewYamlConfigSource("/Go/Bingo-gin/configs/config.yaml")
	app := starters.New(conf)

	app.Start()

	// 测试 ent
	db := base.DbxDataBase()
	client, _ := ent.Open(db)
	u, _ := data.QueryUserById(context.Background(), client, 1)
	fmt.Println(u)
}