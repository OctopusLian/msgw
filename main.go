package main

import (
	"flag"
	"msgw/common/lib"
	"msgw/dao"
	"msgw/router"
	"os"
	"os/signal"
	"syscall"
)

//endpoint dashboard后台管理  server代理服务器
//config ./conf/prod/ 对应配置文件夹

var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *endpoint == "dashboard" {
		lib.InitModule(*config)
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		router.HttpServerStop()
	} else {
		lib.InitModule(*config)
		defer lib.Destroy()
		dao.ServiceManagerHandler.LoadOnce()
		dao.AppManagerHandler.LoadOnce()
	}
}
