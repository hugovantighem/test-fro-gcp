package app

import (
	"context"
	"fmt"
)

func SearchDeletations(ctx context.Context, store DelegationStore, criteria SearchCriteria) ([]Delegation, error) {
	result, err := store.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	return result, nil
}
