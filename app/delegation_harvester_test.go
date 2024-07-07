package app_test

import (
	"context"
	"myproject/app"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestPollDelegations(t *testing.T) {
	ctx := context.Background()
	store := app.NewMockDelegationStore(gomock.NewController(t))
	store.EXPECT().GetLast(ctx).Return(app.Delegation{
		Id: 10,
	}, nil)
	store.EXPECT().Save(ctx, gomock.Any()).Return(nil)
	svc := app.NewMockThezosSvc(gomock.NewController(t))
	svc.EXPECT().GetDelegations(ctx, gomock.Any(), gomock.Any()).Return([]app.DelegationDto{}, nil)
	trigger := NewActivableTrigger()

	stop := app.PollDelegations(ctx, store, svc, trigger)

	trigger.Fire()

	stop <- true
	<-time.After(time.Second)
	assert.True(t, trigger.hasStopped)

}

func TestProcess(t *testing.T) {
	ctx := context.Background()
	store := app.NewMockDelegationStore(gomock.NewController(t))
	store.EXPECT().GetLast(ctx).Return(app.Delegation{
		Id: 10,
	}, nil)
	store.EXPECT().Save(ctx, gomock.Any()).Return(nil)
	svc := app.NewMockThezosSvc(gomock.NewController(t))
	svc.EXPECT().GetDelegations(ctx, gomock.Any(), gomock.Any()).Return([]app.DelegationDto{}, nil)

	err := app.Process(ctx, store, svc)

	assert.NoError(t, err)
}

type ActivableTrigger struct {
	fire       chan time.Time
	hasStopped bool
}

func NewActivableTrigger() *ActivableTrigger {
	return &ActivableTrigger{
		fire:       make(chan time.Time),
		hasStopped: false,
	}
}

func (x *ActivableTrigger) On() <-chan time.Time {
	return x.fire
}

func (x *ActivableTrigger) Fire() {
	x.fire <- time.Now()
}

func (x *ActivableTrigger) Stop() {
	close(x.fire)
	x.hasStopped = true
}
