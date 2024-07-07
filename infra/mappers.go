package infra

import (
	"myproject/api"
	"myproject/app"
	"strconv"
	"time"
)

func ToDelegationDtos(items []app.Delegation) []api.Delegation {
	result := make([]api.Delegation, len(items))
	for idx, _ := range items {
		result[idx] = ToDelegationDto(items[idx])
	}

	return result
}

func ToDelegationDto(item app.Delegation) api.Delegation {
	return api.Delegation{
		Amount:    strconv.Itoa(item.Amount),
		Delegator: item.SenderAddr,
		Level:     strconv.Itoa(item.BlockHeight),
		Timestamp: item.Timestamp.Format(time.RFC3339),
	}
}
