package props

import (
	"github.com/tietang/props/v3/kvs"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
)

// 在应用内的其他地方可以使用 Props() 函数来调用获取所有的额配置
var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	starters.Check(props)
	return props
}


/**
 * 在系统启动的时候，最先是初始化 PropsStarter，会依次执行 starter 接口:
 */
type PropsStarter struct {
	starters.BaseStarter
}


func (s PropsStarter) Init (ctx starters.StarterContext) {
	// 其实这里因为在 boot 的 New() 方法中已经将 conf 配置进入 StarterContext。所以要获取也只是需要从上下文获取就可以了
	//props = ini.NewIniFileConfigSource("config.ini")
	props = ctx.Props()
	//port := props.GetIntDefault("module", "")
	//fmt.Println(port)


	// 在这里初始化的时候还可以动态增加配置，当然写在配置文件也是可以的
	//props.Set("key", "value")
}