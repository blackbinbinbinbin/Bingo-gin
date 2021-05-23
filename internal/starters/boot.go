package starters

import (
	"github.com/tietang/props/v3/kvs"
	"reflect"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"fmt"
)

/**
 * boot.go 文件的作用
 * 主要是利用 BootApplication，启动各个注册 starter 的 init,setup,start 阶段
 * 里边要访问 同一个包下的 starter 内的 starterRegister 数组
 */

// 启动日志
var logger = log.DefaultLogger
func BootLogger() *log.Logger{
	return &logger
}

/**
 * 启动程序的应用管理器
 */
type BootApplication struct {
	starterCtx StarterContext
	conf kvs.ConfigSource
}

//构造系统
func New(conf kvs.ConfigSource) *BootApplication {
	e := &BootApplication{conf: conf, starterCtx: StarterContext{}}
	// 这里通过启动的 main() 内将配置读取后放入 conf 配置，就可以在上下文中传递
	e.starterCtx.SetProps(conf)
	return e
}

// 管理所有的生命周期，启动应用
func (b *BootApplication) Start() {
	//1. 初始化starter
	b.init()
	//2. 安装starter
	b.setup()
	//3. 启动starter
	b.start()
}

/**
 * 初始化starter
 */
func (b *BootApplication) init() {
	log.Info(logger).Log("msg", "Initializing starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		log.Info(logger).Log("msg", fmt.Sprintf("Initializing: PriorityGroup=%d,Priority=%d,type=%s", v.PriorityGroup(), v.Priority(), typ.String()))
		v.Init(b.starterCtx)
	}
}


/**
 * 程序安装
 */
func (b *BootApplication) setup() {
	log.Info(logger).Log("msg", "Setup starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debug(logger).Log("msg", "Setup: "+typ.String())
		v.Setup(b.starterCtx)
	}
}


//程序开始运行，开始接受调用
func (b *BootApplication) start() {
	log.Info(logger).Log("msg", "Starting starters...")
	// 按顺序启动
	allStarters := GetStarters()
	for i, starter := range allStarters {
		if starter.StartBlocking() {
			// 如果阻塞的
			if i == len(allStarters) - 1 {
				// 最后一个直接阻塞启动
				starter.Start(b.starterCtx)
			} else {
				// 如果不是最后一个，使用协程来异步启动
				go starter.Start(b.starterCtx)
			}
		} else {
			// 如果非阻塞的直接启动
			starter.Start(b.starterCtx)
		}
	}
}