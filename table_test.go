package main

import "testing"

func TestSquarePresenter(t *testing.T) {
	cases := map[string]struct {
		square   int
		expected string
	}{
		"fill":     {45, "45 м²"},
		"empty":    {0, "-"},
		"negative": {-4, "-"},
	}

	for n, d := range cases {
		t.Run(n, func(t *testing.T) {
			result := SquarePresenter(d.square)

			if result != d.expected {
				t.Fatalf("Ожидается %s, получено %s", d.expected, result)
			}
		})
	}
}
