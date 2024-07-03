package app

import "context"

func UseCase(ctx context.Context, svc ThezosSvc) error {
	_, err := svc.GetDelegations(ctx, 0)
	return err
}
