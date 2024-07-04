package app

import "context"

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
	Year int
}
