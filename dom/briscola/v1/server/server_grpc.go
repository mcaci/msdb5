package briscola

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	"golang.org/x/net/context"
)

type grpcServer struct {
	points grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints Endpoints) pb.BriscolaServer {
	return &grpcServer{
		points: grpctransport.NewServer(
			endpoints.PointsEndpoint,
			DecodeGRPCPointsRequest,
			EncodeGRPCPointsResponse),
	}
}

func (s *grpcServer) CardPoints(ctx context.Context, r *pb.CardPointsRequest) (*pb.CardPointsResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardPointsResponse), nil
}

func (s *grpcServer) CardCompare(ctx context.Context, r *pb.CardCompareRequest) (*pb.CardCompareResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardCompareResponse), nil
}

func (s *grpcServer) PointCount(ctx context.Context, r *pb.PointCountRequest) (*pb.PointCountResponse, error) {
	_, resp, err := s.points.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PointCountResponse), nil
}

func EncodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(pointsRequest)
	return &pb.CardPointsRequest{CardNumber: uint32(req.CardNumber)}, nil
}

func DecodeGRPCPointsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardPointsRequest)
	return pointsRequest{CardNumber: uint8(req.CardNumber)}, nil
}

func EncodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(pointsResponse)
	return &pb.CardPointsResponse{Points: uint32(res.Points)}, nil
}

func DecodeGRPCPointsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardPointsResponse)
	return pointsResponse{Points: uint8(res.Points), Err: ""}, nil
}
