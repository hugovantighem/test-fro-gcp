package infra

import (
	"context"
	"myproject/api"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {

	return api.GetPing200JSONResponse{
		Ping: "pong",
	}, nil
}

func (Server) GetXtzDelegations(ctx context.Context, request api.GetXtzDelegationsRequestObject) (api.GetXtzDelegationsResponseObject, error) {

	return api.GetXtzDelegations200JSONResponse{}, nil
}
