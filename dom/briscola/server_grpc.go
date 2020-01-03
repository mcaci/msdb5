package briscola

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/pb"
	"golang.org/x/net/context"
)

type grpcServer struct {
	points grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints Endpoints) pb.PointsServer {
	return &grpcServer{
		points: grpctransport.NewServer(
			endpoints.PointsEndpoint,
			DecodeGRPCPointsRequest,
			EncodeGRPCPointsResponse),
	}
}

func (s *grpcServer) Points(ctx context.Context, r *pb.PointsRequest) (*pb.PointsResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PointsResponse), nil
}

func EncodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(pointsRequest)
	return &pb.PointsRequest{Number: uint32(req.Number)}, nil
}

func DecodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PointsRequest)
	return pointsRequest{Number: uint8(req.Number)}, nil
}

func EncodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(pointsResponse)
	return &pb.PointsResponse{Points: uint32(res.Points)}, nil
}

func DecodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.PointsResponse)
	return pointsResponse{Points: uint8(res.Points), Err: ""}, nil
}
