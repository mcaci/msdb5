package briscola

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Service interface {
	CardPoints(ctx context.Context, number uint8) (uint8, error)
	PointCount(ctx context.Context, number uint8) (uint8, error)
	CardCompare(ctx context.Context, number uint8) (uint8, error)
}

type briscolaService struct{}

func (b briscolaService) CardPoints(ctx context.Context, number uint8) (uint8, error) {
	return 15, nil
}

func (b briscolaService) PointCount(ctx context.Context, number uint8) (uint8, error) {
	return 15, nil
}

func (b briscolaService) CardCompare(ctx context.Context, number uint8) (uint8, error) {
	return 15, nil
}

func NewService() Service {
	return briscolaService{}
}

type pointsRequest struct {
	CardNumber uint8 `json:"number"`
}

type pointsResponse struct {
	Points uint8  `json:"points"`
	Err    string `json:"err,omitempty"`
}

func MakePointsEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pointsRequest)
		v, err := srv.CardPoints(ctx, req.CardNumber)
		if err != nil {
			return pointsResponse{v, err.Error()}, nil
		}
		return pointsResponse{v, ""}, nil
	}
}

type Endpoints struct {
	CardPointsEndpoint  endpoint.Endpoint
	PointCountEndpoint  endpoint.Endpoint
	CardCompareEndpoint endpoint.Endpoint
}

func (e Endpoints) CardPoints(ctx context.Context, number uint8) (uint8, error) {
	req := pointsRequest{CardNumber: number}
	resp, err := e.CardPointsEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(pointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

func (e Endpoints) PointCount(ctx context.Context, number uint8) (uint8, error) {
	req := pointsRequest{CardNumber: number}
	resp, err := e.PointCountEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(pointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

func (e Endpoints) CardCompare(ctx context.Context, number uint8) (uint8, error) {
	req := pointsRequest{CardNumber: number}
	resp, err := e.CardCompareEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(pointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}
