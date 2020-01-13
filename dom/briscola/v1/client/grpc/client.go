package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	grpcserv "github.com/mcaci/msdb5/dom/briscola/v1/server/grpc"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) serv.Service {
	var pointsEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "CardPoints",
		grpcserv.EncodeGRPCPointsRequest,
		grpcserv.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()
	var countEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "PointCount",
		grpcserv.EncodeGRPCCountRequest,
		grpcserv.DecodeGRPCCountResponse,
		pb.PointCountResponse{},
	).Endpoint()
	var compareEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "CardCompare",
		grpcserv.EncodeGRPCCompareRequest,
		grpcserv.DecodeGRPCCompareResponse,
		pb.CardCompareResponse{},
	).Endpoint()

	return serv.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
