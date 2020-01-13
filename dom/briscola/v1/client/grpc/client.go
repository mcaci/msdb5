package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	grpcserv "github.com/mcaci/msdb5/dom/briscola/v1/server/grpc"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) serv.Service {
	var pointsEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "Points",
		grpcserv.EncodeGRPCPointsRequest,
		grpcserv.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()
	var countEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "Count",
		grpcserv.EncodeGRPCPointsRequest,
		grpcserv.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()
	var compareEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "Compare",
		grpcserv.EncodeGRPCPointsRequest,
		grpcserv.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()

	return serv.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
