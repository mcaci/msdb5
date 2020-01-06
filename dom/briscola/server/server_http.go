package briscola

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/points", httptransport.NewServer(
		endpoints.PointsEndpoint,
		decodePointsRequest,
		encodePointsResponse,
	))
	return m
}

func decodePointsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pointsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodePointsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
