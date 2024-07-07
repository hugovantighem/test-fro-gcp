package app

import "github.com/sirupsen/logrus"

func ToDeletationModels(items []DelegationDto) []Delegation {
	result := make([]Delegation, len(items))

	for idx := range items {
		logrus.Debugf("mapping id=%d", items[idx].Id)
		result[idx] = ToDeletationModel(items[idx])
	}

	return result
}

func ToDeletationModel(item DelegationDto) Delegation {
	return Delegation{
		Id:          item.Id,
		Amount:      item.Amount,
		SenderAddr:  item.SenderAddr,
		BlockHeight: item.BlockHeight,
		Timestamp:   item.Timestamp,
	}

}
