package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	briscola "github.com/mcaci/msdb5/dom/briscola/v1/server"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) briscola.Service {
	var pointsEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "Points",
		briscola.EncodeGRPCPointsRequest,
		briscola.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()
	var countEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "Count",
		briscola.EncodeGRPCPointsRequest,
		briscola.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()
	var compareEndpoint = grpctransport.NewClient(
		conn, "pb.Briscola", "Compare",
		briscola.EncodeGRPCPointsRequest,
		briscola.DecodeGRPCPointsResponse,
		pb.CardPointsResponse{},
	).Endpoint()

	return briscola.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}
}
