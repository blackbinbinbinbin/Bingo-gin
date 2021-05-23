package init

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
)

// 这里是需要手动注册多个 starter 启动
// 导入包先默认执行引入的包中的 init()
func init() {
	starters.Register(&base.PropsStarter{})
	starters.Register(&base.LogStarter{})
	starters.Register(&base.JsonBaseStarter{})
	starters.Register(&base.DbxDataBaseStarter{})
}