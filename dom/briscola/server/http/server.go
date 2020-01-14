package briscolahttp

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	serv "github.com/mcaci/msdb5/dom/briscola/service"
	"golang.org/x/net/context"
)

func NewHTTPServer(ctx context.Context, endpoints serv.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(
		endpoints.CardPointsEndpoint,
		decodePointsRequest,
		encodePointsResponse,
	))
	m.Handle("/count", httptransport.NewServer(
		endpoints.PointCountEndpoint,
		decodeCountRequest,
		encodeCountResponse,
	))
	m.Handle("/compare", httptransport.NewServer(
		endpoints.CardCompareEndpoint,
		decodeCompareRequest,
		encodeCompareResponse,
	))
	return m
}

func decodePointsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req serv.PointsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodePointsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req serv.CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCompareRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req serv.CompareRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCompareResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
