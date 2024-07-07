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

func ToDataModels(items []app.Delegation) []Delegation {
	result := make([]Delegation, len(items))
	for idx, _ := range items {
		result[idx] = ToDataModel(items[idx])
	}

	return result
}

func ToDataModel(item app.Delegation) Delegation {
	return Delegation{
		Id:          item.Id,
		Amount:      item.Amount,
		SenderAddr:  item.SenderAddr,
		BlockHeight: item.BlockHeight,
		Timestamp:   item.Timestamp,
		Year:        item.Timestamp.Year(),
	}
}

func FromDataModels(items []Delegation) []app.Delegation {
	result := make([]app.Delegation, len(items))
	for idx := range items {
		result[idx] = FromDataModel(items[idx])
	}

	return result
}

func FromDataModel(item Delegation) app.Delegation {
	return app.Delegation{
		Id:          item.Id,
		Amount:      item.Amount,
		SenderAddr:  item.SenderAddr,
		BlockHeight: item.BlockHeight,
		Timestamp:   item.Timestamp,
	}
}
