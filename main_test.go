package main_test

import (
	"myproject/infra"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestApplication(t *testing.T) {
	addr := "0.0.0.0:8080"
	stop := infra.RunApplication(addr)
	defer stop()

	<-time.After(time.Second)
	res, err := http.Get("http://" + addr + "/ping")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
