package app

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

const noLastId = -1

// PollDelegations fetches delegation from external service whenever the trigger fires.
//
// delegation are fetched based on the highest fetched id.
//
// returns a channel to complete gracefully by stopping the trigger and stop polling delegations.
func PollDelegations(ctx context.Context, store DelegationStore, tzsSvc ThezosSvc, trigger Trigger) chan<- bool {

	result := make(chan bool)
	fire := trigger.On()

	go func() {
		for {
			select {
			case <-result:
				logrus.Info("stop polling")
				trigger.Stop()
				return
			case <-fire:
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

// process get the last id stored.
//
// fetches from external service for id > found id or from the beginning if not previous entry is found.
//
// saves the ne entries.
//
// returns an error if any.
//
// NOTE: should return Retryable and NonRetryable errors.
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

	err = store.Save(ctx, ToDomainModels(items))
	if err != nil {
		return err
	}

	return nil
}
