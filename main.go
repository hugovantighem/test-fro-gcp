package main

import (
	"log"
	"myproject/infra"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	db, err := infra.RunMigrateScripts()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	stop := infra.RunApplication("0.0.0.0:8080")
	defer stop()
	<-done

}
