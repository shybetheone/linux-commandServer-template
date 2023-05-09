package inits

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"linux-service-template/cmd"
	"os"
)

var (
	addr string
	port string
	vp   string
)

func NewUrfaveCli() {
	app := &cli.App{
		Name:                   "vm-vm-proxy",
		Usage:                  "a vm to bm proxy",
		Version:                "v0.0.1",
		UseShortOptionHandling: true,
		Flags:                  cmd.GlobalOptions,
		// Before 在任意命令执行前执行，这里用来处理全局选项
		Before: cmd.LoadGlobalOptions,
		// 同理，也可以定义 After 来执行收尾操作

		Action: func(c *cli.Context) error {
			addr = c.String("addr")
			port = c.String("port")
			vp = c.String("vp")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil && err != cmd.ErrPrintAndExit {
		log.Fatal(err)
	}
	if port == "" || addr == "" || vp == "" {
		os.Exit(1)
		return
	}

}
