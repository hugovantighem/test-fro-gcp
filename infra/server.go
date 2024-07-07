package infra

import (
	"context"
	"myproject/api"
	"myproject/app"
)

type Server struct {
	store app.DelegationStore
}

func NewServer(store app.DelegationStore) *Server {
	return &Server{
		store: store,
	}
}

func (x *Server) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {

	return api.GetPing200JSONResponse{
		Ping: "pong",
	}, nil
}

func (x *Server) GetXtzDelegations(ctx context.Context, request api.GetXtzDelegationsRequestObject) (api.GetXtzDelegationsResponseObject, error) {
	criteria := app.SearchCriteria{}

	if request.Params.Year != nil {
		year := int(*request.Params.Year)
		criteria.Year = &year
	}

	items, err := app.SearchDeletations(ctx, x.store, criteria)
	if err != nil {
		return api.GetXtzDelegations500JSONResponse{Message: err.Error()}, nil
	}

	return api.GetXtzDelegations200JSONResponse{
		Data: ToDelegationDtos(items),
	}, nil
}
