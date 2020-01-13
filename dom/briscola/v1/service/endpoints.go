package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type PointsRequest struct {
	CardNumber uint32 `json:"number"`
}

type PointsResponse struct {
	Points uint32 `json:"points"`
	Err    string `json:"err,omitempty"`
}

func MakePointsEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PointsRequest)
		v, err := srv.CardPoints(ctx, req.CardNumber)
		if err != nil {
			return PointsResponse{v, err.Error()}, nil
		}
		return PointsResponse{v, ""}, nil
	}
}

type Endpoints struct {
	CardPointsEndpoint  endpoint.Endpoint
	PointCountEndpoint  endpoint.Endpoint
	CardCompareEndpoint endpoint.Endpoint
}

func (e Endpoints) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	req := PointsRequest{CardNumber: number}
	resp, err := e.CardPointsEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(PointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

func (e Endpoints) PointCount(ctx context.Context, number uint32) (uint32, error) {
	req := PointsRequest{CardNumber: number}
	resp, err := e.PointCountEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(PointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}

func (e Endpoints) CardCompare(ctx context.Context, number uint32) (uint32, error) {
	req := PointsRequest{CardNumber: number}
	resp, err := e.CardCompareEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	pointsResp := resp.(PointsResponse)
	if pointsResp.Err != "" {
		return 0, errors.New(pointsResp.Err)
	}
	return pointsResp.Points, nil
}
