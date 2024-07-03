package app

import "context"

type DelegationStore interface {
	// -- Read
	// Search for all delegations, may be filtered by year.
	// Note: pagination is not implemented
	Search(ctx context.Context, criteria SearchCriteria) ([]Delegation, error)
	// Retrieve the last Delegation (highest id).
	GetLast(ctx context.Context) (Delegation, error)

	// -- Write
	Save(ctx context.Context, item Delegation) error
}

type SearchCriteria struct {
	Year int
}
