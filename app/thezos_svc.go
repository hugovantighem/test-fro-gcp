package app

import (
	"context"
	"time"
)

//go:generate mockgen -package=app -source=thezos_svc.go -destination=thezos_svc.mock.go
type ThezosSvc interface {
	GetDelegations(ctx context.Context, id int) ([]DelegationDto, error)
}

type DelegationDto struct {
	Id          int
	Amount      int
	Timestamp   time.Time
	SenderAddr  string
	BlockHeight int
}
