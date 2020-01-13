package briscola

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Service interface {
	Points(ctx context.Context, number uint8) (uint8, error)
}

type pointsService struct{}

func (p pointsService) Points(ctx context.Context, number uint8) (uint8, error) {
	return 15, nil
}

func NewService() Service {
	return pointsService{}
}

type PointsRequest struct {
	Number uint8 `json:"number"`
}

type PointsResponse struct {
	Points uint8  `json:"points"`
	Err    string `json:"err,omitempty"`
}

func MakePointsEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PointsRequest)
		v, err := srv.Points(ctx, req.Number)
		if err != nil {
			return PointsResponse{v, err.Error()}, nil
		}
		return PointsResponse{v, ""}, nil
	}
}

type Endpoints struct {
	PointsEndpoint endpoint.Endpoint
}

func (e Endpoints) Points(ctx context.Context, number uint8) (uint8, error) {
	req := PointsRequest{Number: number}
	resp, err := e.PointsEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(PointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}
