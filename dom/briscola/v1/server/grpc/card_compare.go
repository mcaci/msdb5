package briscola

import (
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	"golang.org/x/net/context"
)

func (s *grpcServer) CardCompare(ctx context.Context, r *pb.CardCompareRequest) (*pb.CardCompareResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardCompareResponse), nil
}

func EncodeGRPCCompareRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(serv.PointsRequest)
	return &pb.CardPointsRequest{CardNumber: req.CardNumber}, nil
}

func DecodeGRPCCompareRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return serv.PointsRequest{CardNumber: req.CardNumber}, nil
}

func EncodeGRPCCompareResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(serv.PointsResponse)
	return &pb.CardPointsResponse{Points: res.Points}, nil
}

func DecodeGRPCCompareResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return serv.PointsResponse{Points: res.Points, Err: ""}, nil
}
