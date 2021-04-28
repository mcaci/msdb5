package srv

import (
	"context"
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
)

type server struct {
	pb.UnimplementedBriscolaServer
}

func NewServer() *server { return &server{} }

func (s *server) Points(ctx context.Context, n *pb.NumberReq) (*pb.PointsRes, error) {
	c, err := card.FromID(uint8(n.Number))
	if err != nil {
		return &pb.PointsRes{}, err
	}
	return &pb.PointsRes{Points: uint32(briscola.Points(c))}, nil
}

func (s *server) Score(ctx context.Context, cs *pb.Cards) (*pb.PointsRes, error) {
	crds, err := toCards(cs)
	if err != nil {
		return nil, fmt.Errorf("Score: error in the conversion of the input: %w", err)
	}
	return &pb.PointsRes{Points: uint32(briscola.Score(*crds))}, nil
}

func (s *server) Winner(ctx context.Context, b *pb.Board) (*pb.Index, error) {
	crds, err := toCards(b.Cards)
	if err != nil {
		return nil, fmt.Errorf("Winner: error in the conversion of the input: %w", err)
	}
	id := briscola.Winner(*crds, card.Seed(uint8(b.Briscola)))
	return &pb.Index{Id: uint32(id)}, nil
}
