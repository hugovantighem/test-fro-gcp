package infra

import (
	"context"
	"errors"
	"myproject/api"
	"myproject/app"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ShutDownFuncs []func()

func (x ShutDownFuncs) Shutdown() {
	for i := len(x) - 1; i >= 0; i-- {
		x[i]()
	}
}

func RunApplication(conf Config) func() {
	var store app.DelegationStore

	var shutdown ShutDownFuncs = []func(){}

	if conf.DbConnString == "" {
		store = NewInMemoryDelegationStorage()
	} else {
		db := InitDB(conf)
		shutdown = append(shutdown, func() {
			db, err := db.DB()
			if err != nil {
				return
			}
			db.Close()
		})
		store = NewPgStorage(db)
	}

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := NewServer(store)

	r := gin.Default()

	handler := api.NewStrictHandler(server, nil)
	api.RegisterHandlers(r, handler)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    conf.ServerAddr,
	}

	// And we serve HTTP until the world ends.
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("error: %s\n", err)
		}
	}()

	ctx := context.Background()
	trigger := app.NewTickerTrigger(5 * time.Second)

	quit := app.PollDelegations(ctx,
		store,
		NewTzktClient(http.DefaultClient, conf.ThezosApiAddr),
		trigger,
	)
	shutdown = append(shutdown, func() { quit <- true })

	return func() {

		logrus.Println("Shutting down server...")
		shutdown.Shutdown()

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			logrus.Fatalf("Error while shutting down Server: %s. Initiating force shutdown...", err.Error())
		} else {
			logrus.Info("Server exiting")
		}
	}
}
