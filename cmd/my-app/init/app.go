package init

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/props"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/log"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/json"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/db"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/http"
)

// 这里是需要手动注册多个 starter 启动
// 导入包先默认执行引入的包中的 init()
func init() {
	starters.Register(&props.PropsStarter{})
	starters.Register(&log.LogStarter{})
	starters.Register(&json.JsonBaseStarter{})
	starters.Register(&db.DbxDataBaseStarter{})
	starters.Register(&http.GinServerStart{})
}