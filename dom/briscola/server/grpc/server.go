package briscola

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/pb"
	serv "github.com/mcaci/msdb5/dom/briscola/service"
	"golang.org/x/net/context"
)

type grpcServer struct {
	points  grpctransport.Handler
	count   grpctransport.Handler
	compare grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints serv.Endpoints) pb.BriscolaServer {
	return &grpcServer{
		points: grpctransport.NewServer(
			endpoints.CardPointsEndpoint,
			DecodeGRPCPointsRequest,
			EncodeGRPCPointsResponse),
		count: grpctransport.NewServer(
			endpoints.PointCountEndpoint,
			DecodeGRPCCountRequest,
			EncodeGRPCCountResponse),
		compare: grpctransport.NewServer(
			endpoints.CardCompareEndpoint,
			DecodeGRPCCompareRequest,
			EncodeGRPCCompareResponse),
	}
}
