package briscola

import (
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	"golang.org/x/net/context"
)

func (s *grpcServer) CardPoints(ctx context.Context, r *pb.CardPointsRequest) (*pb.CardPointsResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardPointsResponse), nil
}

func EncodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(serv.PointsRequest)
	return &pb.CardPointsRequest{CardNumber: uint32(req.CardNumber)}, nil
}

func DecodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return serv.PointsRequest{CardNumber: uint8(req.CardNumber)}, nil
}

func EncodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(serv.PointsResponse)
	return &pb.CardPointsResponse{Points: uint32(res.Points)}, nil
}

func DecodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return serv.PointsResponse{Points: uint8(res.Points), Err: ""}, nil
}
