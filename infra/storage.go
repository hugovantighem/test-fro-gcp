package infra

import (
	"cmp"
	"context"
	"myproject/app"
	"slices"
	"sync"
)

type InMemoryDelegationStorage struct {
	mu    sync.Mutex
	items []app.Delegation
}

func NewInMemoryDelegationStorage() InMemoryDelegationStorage {
	return InMemoryDelegationStorage{
		mu:    sync.Mutex{},
		items: make([]app.Delegation, 0),
	}
}

func (x *InMemoryDelegationStorage) Search(ctx context.Context, criteria app.SearchCriteria) ([]app.Delegation, error) {
	x.mu.Lock()
	defer x.mu.Unlock()
	slices.SortFunc(x.items, func(a, b app.Delegation) int { return cmp.Compare(a.Id, b.Id) })
	return x.items, nil
}
func (x *InMemoryDelegationStorage) GetLast(ctx context.Context) (app.Delegation, error) {
	x.mu.Lock()
	defer x.mu.Unlock()
	return slices.MaxFunc(x.items, func(a, b app.Delegation) int { return cmp.Compare(a.Id, b.Id) }), nil
}
func (x *InMemoryDelegationStorage) Save(ctx context.Context, item app.Delegation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	x.items = append(x.items, item)
	return nil
}
