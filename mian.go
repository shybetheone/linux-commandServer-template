package main

import (
	"linux-service-template/inits"

	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan os.Signal)
	inits.LogInit()
	inits.NewUrfaveCli()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server exiting")

}
