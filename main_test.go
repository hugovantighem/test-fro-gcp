package main_test

import (
	"myproject/infra"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestApplication(t *testing.T) {
	conf := infra.Config{
		ServerAddr: "0.0.0.0:8080",
	}
	stop := infra.RunApplication(conf)
	defer stop()

	<-time.After(time.Second)
	res, err := http.Get("http://" + conf.ServerAddr + "/ping")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
