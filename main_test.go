package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "zero size", size: 0},
		{name: "negative size", size: -5},
		{name: "one element", size: 1},
		{name: "many elements", size: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			wantLen := tt.size
			if wantLen < 0 {
				wantLen = 0
			}
			if len(got) != wantLen {
				t.Fatalf("len = %d, want %d", len(got), wantLen)
			}
			for i, value := range got {
				if value <= 0 {
					t.Errorf("element %d = %d, want positive value", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "max at start", data: []int{9, 1, 2}, want: 9},
		{name: "max in middle", data: []int{1, 9, 2}, want: 9},
		{name: "max at end", data: []int{1, 2, 9}, want: 9},
		{name: "all equal", data: []int{4, 4, 4}, want: 4},
		{name: "negative values", data: []int{-3, -1, -7}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{5}, want: 5},
		{name: "fewer elements than chunks", data: []int{2, 8, 3}, want: 8},
		{name: "length not divisible by chunks", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, want: 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunksMatchesMaximum проверяет слайсы разной длины с максимумом в каждой позиции,
// чтобы покрыть все границы частей и обработку остатка.
func TestMaxChunksMatchesMaximum(t *testing.T) {
	for length := 1; length <= 40; length++ {
		for pos := 0; pos < length; pos++ {
			data := make([]int, length)
			for i := range data {
				data[i] = 1
			}
			data[pos] = 100

			if got := maxChunks(data); got != 100 {
				t.Errorf("length %d, max at %d: maxChunks = %d, want 100", length, pos, got)
			}
		}
	}
}
