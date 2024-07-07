package app_test

import (
	"myproject/app"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToDeletationModel(t *testing.T) {
	ts, err := time.Parse(time.RFC3339, "2019-08-24T14:15:22Z")
	require.NoError(t, err)
	item := app.DelegationDto{
		Id:         5084545024,
		Amount:     30000000000,
		SenderAddr: "tz3WuwwhU8kVReq8f8TJZ1g14mHBmsr8meGn",
		Timestamp:  ts,
	}

	result := app.ToDeletationModel(item)

	assert.Equal(t, 5084545024, result.Id)
	assert.Equal(t, 30000000000, result.Amount)
	assert.Equal(t, "tz3WuwwhU8kVReq8f8TJZ1g14mHBmsr8meGn", result.SenderAddr)
	assert.Equal(t, ts, result.Timestamp)
	assert.Equal(t, "30000000000", strconv.Itoa(result.Amount))
}
