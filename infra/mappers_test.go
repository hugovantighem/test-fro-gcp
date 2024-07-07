package infra_test

import (
	"myproject/app"
	"myproject/infra"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToDelegationDto(t *testing.T) {
	item := app.Delegation{
		Id:          1,
		Amount:      100,
		SenderAddr:  "my-addr",
		BlockHeight: 175,
		Timestamp:   time.Now(),
	}

	result := infra.ToDelegationDto(item)

	assert.Equal(t, strconv.Itoa(item.Amount), result.Amount)
	assert.Equal(t, strconv.Itoa(item.BlockHeight), result.Level)
	assert.Equal(t, item.SenderAddr, result.Delegator)
	assert.Equal(t, item.Timestamp.Format(time.RFC3339), result.Timestamp)

}

func TestToDataModel(t *testing.T) {
	ts, err := time.Parse(time.RFC3339, "2019-08-24T14:15:22Z")
	require.NoError(t, err)

	item := app.Delegation{
		Id:          1,
		Amount:      100,
		SenderAddr:  "my-addr",
		BlockHeight: 175,
		Timestamp:   ts,
	}

	result := infra.ToDataModel(item)

	assert.Equal(t, item.Id, result.Id)
	assert.Equal(t, item.Amount, result.Amount)
	assert.Equal(t, item.SenderAddr, result.SenderAddr)
	assert.Equal(t, item.BlockHeight, result.BlockHeight)
	assert.Equal(t, item.Timestamp, result.Timestamp)
	assert.Equal(t, 2019, result.Year)

}

func TestFromDataModel(t *testing.T) {
	item := infra.Delegation{
		Id:          1,
		Amount:      100,
		SenderAddr:  "my-addr",
		BlockHeight: 175,
		Timestamp:   time.Now(),
	}

	result := infra.FromDataModel(item)

	assert.Equal(t, item.Id, result.Id)
	assert.Equal(t, item.Amount, result.Amount)
	assert.Equal(t, item.SenderAddr, result.SenderAddr)
	assert.Equal(t, item.BlockHeight, result.BlockHeight)
	assert.Equal(t, item.Timestamp, result.Timestamp)

}
