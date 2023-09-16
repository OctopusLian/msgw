package main

import (
	"msgw/common/lib"
	"msgw/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lib.InitModule("./conf/dev/")
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
