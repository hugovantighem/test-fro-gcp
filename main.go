package main

import (
	"myproject/infra"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	stop := infra.RunApplication("0.0.0.0:8080")

	<-done
	stop()

}
