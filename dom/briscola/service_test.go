package briscola

import (
	"context"
	"testing"
)

func TestPointsService(t *testing.T) {
	srv := NewService()
	ctx := context.Background()
	points, err := srv.Points(ctx, 1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(points)
}
