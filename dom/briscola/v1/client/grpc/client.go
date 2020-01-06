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

	return briscola.Endpoints{
		PointsEndpoint: pointsEndpoint,
	}
}
