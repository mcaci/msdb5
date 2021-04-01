package briscola5

import (
	"testing"
)

var initialValue AuctionScore

func TestAuctionCases(t *testing.T) {
	testcases := map[string]struct {
		current, proposed AuctionScore
		expected          cmpInfo
	}{
		"Current and proposed less than min -> MIN_SCORE":         {0, 1, LT_MIN_SCORE},
		"Proposed less than current but more than min -> current": {76, 65, LE_ACTUAL},
		"Normal case -> proposed":                                 {61, 65, GT_ACTUAL},
		"Proposed higher or equal than max -> MAX_SCORE":          {110, 121, GE_MAX_SCORE},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if info := Cmp(tc.current, tc.proposed); info != tc.expected {
				t.Fatalf("Compare info's value should be %d but is %d", tc.expected, info)
			}
		})
	}
}

func TestCheckAndUpdate_OK(t *testing.T) {
	value := AuctionScore(80)
	if Cmp(value, AuctionScore(100)) != 0 {
		t.Fatal("Unexpected check return value")
	}
}

func TestCheckAndUpdate_Fold(t *testing.T) {
	value := AuctionScore(80)
	if Cmp(value, AuctionScore(61)) >= 0 {
		t.Fatal("Unexpected check return value")
	}
}

func TestAuctionValues(t *testing.T) {
	testcases := map[string]struct {
		current, proposed, expected AuctionScore
	}{
		"Current and proposed less than min -> MIN_SCORE":         {0, 1, MIN_SCORE},
		"Proposed less than current but more than min -> current": {76, 65, 76},
		"Normal case -> proposed":                                 {61, 65, 65},
		"Proposed higher or equal than max -> MAX_SCORE":          {110, 121, MAX_SCORE},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if s := CmpAndSet(tc.current, tc.proposed); s != tc.expected {
				t.Fatalf("Auction score should be set at %d but is %d", tc.expected, s)
			}
		})
	}
}
