package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endp "github.com/mcaci/msdb5/dom/briscola/endpoint"
	"github.com/mcaci/msdb5/dom/briscola/pb"
	grpcserv "github.com/mcaci/msdb5/dom/briscola/server/grpc"
	serv "github.com/mcaci/msdb5/dom/briscola/service"
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

	return endp.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
