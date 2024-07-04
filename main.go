package main

import (
	"log"
	"myproject/infra"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := infra.RunMigrateScripts()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	stop := infra.RunApplication("0.0.0.0:8080")

	<-done
	stop()

}
