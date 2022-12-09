package main

import "testing"

func TestPointDiff(t *testing.T) {
	testCases := []struct{ p1, p2, diff int }{
		{0, 1, -1},
		{1, 0, 1},
		{0, -1, 1},
		{-1, 0, -1},
		{5, 3, 2},
		{3, 5, -2},
		{-5, -2, -3},
		{-2, -5, 3},
		{-5, 2, 7},
		{2, -5, -7},
	}

	for _, tc := range testCases {
		diff := pointDiff(tc.p1, tc.p2)
		if diff != tc.diff {
			t.Fatalf("(%d, %d) expected: %d, calculated %d", tc.p1, tc.p2, tc.diff, diff)
		}
	}
}
