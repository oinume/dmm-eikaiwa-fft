package grpc_server

import (
	api_v1 "github.com/oinume/lekcije/proto-gen/go/proto/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type apiV1Server struct{}

func RegisterAPIV1Server(server *grpc.Server) {
	api_v1.RegisterAPIServer(server, &apiV1Server{})
}

func (s *apiV1Server) GetSettingEmail(
	ctx context.Context, in *api_v1.GetSettingEmailRequest,
) (*api_v1.GetSettingEmailResponse, error) {
	// TODO: implement
	return &api_v1.GetSettingEmailResponse{Email:"a@foo.com"}, nil
}
