package app

import (
	"context"
	"fmt"
)

//go:generate mockgen -package=app -source=storage.go -destination=storage.mock.go
type DelegationStore interface {
	// -- Read

	// Search for all delegations, may be filtered by year.
	// Note: pagination is not implemented
	Search(ctx context.Context, criteria SearchCriteria) ([]Delegation, error)

	// Retrieves the last Delegation (highest id).
	// returns the the found element if any.
	// returns a ErrNotFound if there is no element in the collection.
	// returns a ErrTechnical if any.
	GetLast(ctx context.Context) (Delegation, error)

	// -- Write

	Save(ctx context.Context, items []Delegation) error
}

type SearchCriteria struct {
	Year *int
}

func (x SearchCriteria) String() string {
	if x.Year == nil {
		return "no search criteria"
	}

	return fmt.Sprintf("year=%d", *x.Year)
}
