package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/pb"
	briscola "github.com/mcaci/msdb5/dom/briscola/server"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) briscola.Service {
	var pointsEndpoint = grpctransport.NewClient(
		conn, "pb.Points", "Points",
		briscola.EncodeGRPCPointsRequest,
		briscola.DecodeGRPCPointsResponse,
		pb.PointsResponse{},
	).Endpoint()

	return briscola.Endpoints{
		PointsEndpoint: pointsEndpoint,
	}
}
