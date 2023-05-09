package cmd

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var addr string
var port string
var vp string

// GlobalOptions 将全局选项暴露给 main 包
var GlobalOptions = []cli.Flag{
	&cli.StringFlag{
		Name:     "port",
		Usage:    "a tcp port you have to write",
		Required: true,
	},
	&cli.StringFlag{
		Name:     "addr",
		Usage:    "a tcp address you have to write",
		Required: true,
	},
	&cli.StringFlag{
		Name:     "vp",
		Usage:    "a vsock port you have to write",
		Required: true,
	},
	// 省略 man-path, info-path, paginate, no-pager...
	// more ...
}

// ErrPrintAndExit 表示遇到需要打印信息并提前退出的情形，不需要打印错误信息
var ErrPrintAndExit = errors.New("print and exit")

// LoadGlobalOptions 加载全局选项
var LoadGlobalOptions = func(ctx *cli.Context) error {
	// 并非实际实现，所以遇到对应的参数只是输出信息，方便观察
	// 全局选项既可以在这里读取并设置全局状态（如有）
	// 也可以在具体实现处再通过 ctx 读取（参考 add）
	if ctx.IsSet("port") {
		log.Info("started dial tcp port")
	} else {
		log.Println("port flag must be set")
	}
	// 省略 exec-path ...
	if ctx.IsSet("addr") {
		log.Info("started dial tcp address")
	} else {
		log.Println("addr flag must be set")
	}
	// 省略 man-path, info-path ...
	if ctx.IsSet("vp") {
		log.Info("started listen vsock port")
	} else {
		log.Println("vp flag must be set")
	}
	return nil
}

// 子命令分组
const (
	cmdGroupStart = "start a working area"
	cmdGroupWork  = "work on current change"
	// ...
)
