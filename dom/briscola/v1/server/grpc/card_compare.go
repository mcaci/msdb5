package briscola

import (
	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	"golang.org/x/net/context"
)

func (s *grpcServer) CardCompare(ctx context.Context, r *pb.CardCompareRequest) (*pb.CardCompareResponse, error) {
	_, resp, err := s.compare.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CardCompareResponse), nil
}

func EncodeGRPCCompareRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(serv.CompareRequest)
	first := &pb.ItalianCard{Number: req.FirstCardNumber, Seed: pb.Seed(req.FirstCardSeed)}
	second := &pb.ItalianCard{Number: req.SecondCardNumber, Seed: pb.Seed(req.SecondCardSeed)}
	return &pb.CardCompareRequest{FirstCard: first, SecondCard: second, Briscola: pb.Seed(req.BriscolaSeed)}, nil
}

func DecodeGRPCCompareRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CardCompareRequest)
	return serv.CompareRequest{
		FirstCardNumber:  req.FirstCard.GetNumber(),
		FirstCardSeed:    uint32(req.FirstCard.GetSeed()),
		SecondCardNumber: req.SecondCard.GetNumber(),
		SecondCardSeed:   uint32(req.SecondCard.GetSeed()),
		BriscolaSeed:     uint32(req.Briscola)}, nil
}

func EncodeGRPCCompareResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(serv.CompareResponse)
	return &pb.CardCompareResponse{SecondCardWinsOverFirstOne: res.SecondCardWins}, nil
}

func DecodeGRPCCompareResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.CardCompareResponse)
	return serv.CompareResponse{SecondCardWins: res.SecondCardWinsOverFirstOne, Err: ""}, nil
}
