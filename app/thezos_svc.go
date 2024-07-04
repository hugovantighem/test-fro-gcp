package app

import (
	"context"
	"time"
)

//go:generate mockgen -package=app -source=thezos_svc.go -destination=thezos_svc.mock.go
type ThezosSvc interface {
	// fetches deletations having id > the given id. Expects id == -1 to fetch with no filter on the id.
	GetDelegations(ctx context.Context, id int, resultLimit int) ([]DelegationDto, error)
}

type DelegationDto struct {
	Id          int
	Amount      int
	Timestamp   time.Time
	SenderAddr  string
	BlockHeight int
}
