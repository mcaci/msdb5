package briscola

import "context"

type Service interface {
	Points(ctx context.Context, number uint8) (uint8, error)
}

type pointsService struct{}

func (p pointsService) Points(ctx context.Context, number uint8) (uint8, error) {
	return 0, nil
}

func NewService() Service {
	return pointsService{}
}
