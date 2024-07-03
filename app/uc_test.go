package app_test

import (
	"context"
	"fmt"
	"myproject/app"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUC(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()
		svc := app.NewMockThezosSvc(gomock.NewController(t))
		svc.EXPECT().GetDelegations(ctx, 0).Return(app.Delegation{}, fmt.Errorf("error"))

		err := app.UseCase(ctx, svc)
		assert.Error(t, err)
	})
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		svc := app.NewMockThezosSvc(gomock.NewController(t))
		svc.EXPECT().GetDelegations(ctx, 0).Return(app.Delegation{}, nil)

		err := app.UseCase(ctx, svc)
		assert.NoError(t, err)
	})

}
