package serv

import (
	"context"
)

type Service interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
	PointCount(ctx context.Context, number []uint32) (uint32, error)
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}

type briscolaService struct{}

func (b briscolaService) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	return 15, nil
}

func (b briscolaService) PointCount(ctx context.Context, number []uint32) (uint32, error) {
	return 15, nil
}

func (b briscolaService) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	return true, nil
}

func NewService() Service {
	return briscolaService{}
}
