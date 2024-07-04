package main

import (
	"context"
	"log"
	"myproject/infra"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	var conf Config
	if err := envconfig.Process(ctx, &conf); err != nil {
		log.Fatal(err)
	}

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
