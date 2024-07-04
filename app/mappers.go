package app

import "github.com/sirupsen/logrus"

func ToDeletationModel(items []DelegationDto) []Delegation {
	result := make([]Delegation, len(items))

	for idx := range items {
		logrus.Debugf("mapping id=%d", items[idx].Id)
		result = append(result, Delegation{
			Id:          items[idx].Id,
			Amount:      items[idx].Amount,
			SenderAddr:  items[idx].SenderAddr,
			BlockHeight: items[idx].BlockHeight,
			Timestamp:   items[idx].Timestamp,
			Year:        items[idx].Timestamp.Year(),
		})
	}

	return result
}
