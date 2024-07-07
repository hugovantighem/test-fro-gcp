package app

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

const noLastId = -1

func PollDelegations(ctx context.Context, store DelegationStore, tzsSvc ThezosSvc) chan<- bool {

	result := make(chan bool)
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-result:
				logrus.Info("stop polling")
				ticker.Stop()
				return
			case <-ticker.C:
				err := process(ctx, store, tzsSvc)
				if err != nil {
					logrus.Error(err)
					continue
				}
			}
		}
	}()

	return result
}

func process(ctx context.Context, store DelegationStore, tzsSvc ThezosSvc) error {
	item, err := store.GetLast(ctx)

	var lastId int
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			lastId = noLastId
		} else {
			return err
		}
	} else {
		lastId = item.Id
	}

	logrus.Debugf("polling... from id=%d", lastId)

	items, err := tzsSvc.GetDelegations(ctx, lastId, 5)
	if err != nil {
		return err
	}

	err = store.Save(ctx, ToDeletationModel(items))
	if err != nil {
		return err
	}

	return nil
}
