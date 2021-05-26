package main

import (
	"github.com/tietang/props/v3/yam"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	_ "github.com/blackbinbinbinbin/Bingo-gin/cmd/my-app/init"
)

func main() {
	// 加载配置文件
	conf := yam.NewYamlConfigSource("/Go/Bingo-gin/configs/config.yaml")
	app := starters.New(conf)

	go app.Stop()

	app.Start()
}